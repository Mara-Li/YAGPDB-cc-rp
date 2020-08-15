{{/* CONFIGURATION AREA STARTS */}}
{{$Suggestion_Channel:=741805466761560134}}
{{$Logging_Channel:=741808846473003088}}
{{$Implemented_Channel:=701792582362988615}}
{{$Approved_Channel:= 741805466761560134}}
{{$Mod_Roles:=cslice 453167633068457984 }}{{/* No need to add Admin roles. They are automatically detected given Yag has right Perms */}}
{{$Cooldown:=0}}{{/* Can be set to 0 for no cooldown */}}
{{$Upvote:="this:715708401816043520"}}
{{$Downvote:=":moins:715708440466554921"}}
{{/* CONFIGURATION AREA ENDS */}}
{{$globalDict:=dict "chans" (dict $Suggestion_Channel true $Approved_Channel true $Implemented_Channel true) "msg" .nil}}
{{$Prefix:=index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}
{{$Escaped_Prefix:=reReplace `[\.\[\]\-\?\!\\\*\{\}\(\)\|]` $Prefix `\${0}`}}
{{$error:=""}}
{{$Syntax:=""}}
{{$IS_Mod:=false}}
{{if in (slice (exec "viewperms") (add 25 (len .User.Username))) `Administrator`}}{{$IS_Mod =true}}{{else}}{{range $Mod_Roles}}{{if in $.Member.Roles .}}{{$IS_Mod =true}}{{end}}{{end}}{{end}}
{{$Attachments:=""}}{{$Img_Set:=false}}
{{if not (reFind `\$` .Message.Content)}}{{$error ="Le regex ne correspond pas au prefix"}}{{$Syntax =`Prefix/Yag Mention <Command> <Args>`}}{{else}}
{{if reFind `(?i)\bsuggest(ion)?\b` .Cmd}}
	{{$Syntax =print .Cmd " <Suggestion_Here>"}}
	{{$col:=16777215}}{{$pos:=0}}{{$r:=.Member.Roles}}{{range .Guild.Roles}}{{if and (in $r .ID) (.Color) (lt $pos .Position)}}{{$pos =.Position}}{{$col =.Color}}{{end}}{{end}}
	{{if or .StrippedMsg .Message.Attachments}}
		{{with (dbGet .User.ID "suggestCld")}}
			{{$error =print "Cooldown de : " (humanizeDurationSeconds (.ExpiresAt.Sub currentTime)) " pour éviter le spamm."}}
		{{else}}
			{{if not $IS_Mod}}{{if $Cooldown}}{{dbSetExpire .User.ID "suggestCld" "cooldown" $Cooldown}}{{end}}{{end}}
	{{$c:= .StrippedMsg}}
	{{$title := ""}}
	{{if reFind `(\#\S*)` $c}}
		{{$title = reFind `(\#\S*)` $c}}
		{{$c = joinStr " " (split $c $title)}}
		{{$c = joinStr  " " (split $c "#")}}
		{{$title = joinStr " " (print "Suggestion #" (dbIncr 0 "suggestions-count" 1) " — " (split $title "#"))}}
	{{else}}
		{{$title = (print "Suggestion #" (dbIncr 0 "suggestions-count" 1))}}
	{{end}}
			{{$embed :=	sdict
						"title" $title
						"description" $c
						"color" $col
						"author" (sdict "name" (str $.User) "icon_url" ($.User.AvatarURL "512"))
						"timestamp"  currentTime
						"footer" (sdict "text" (print "Author ID - " $.User.ID))
			}}


			{{range .Message.Attachments}}{{if and (not $Img_Set) .Width}}{{$Img_Set =true}}{{$embed.Set "image" (sdict "url" .ProxyURL)}}{{else}}{{$Attachments =print $Attachments "\n[" .Filename "](" .ProxyURL ")"}}{{end}}{{end}}
			{{if $Attachments}}{{$embed.Set "description" (print $embed.description "\n\n**__Attachments:__**" $Attachments)}}{{end}}
			{{$ID:=sendMessageRetID $Suggestion_Channel (complexMessage "content" "@here : Vous avez une nouvelle suggestion ! \n *Vous pouvez en discuter dans le channel <#741808846473003088>*" "embed" $embed)}}
			{{addMessageReactions $Suggestion_Channel $ID $Upvote $Downvote}}
			{{addReactions $Upvote}}
		{{end}}
	{{else}}
		{{$error ="Arguments insuffisants"}}
	{{end}}

{{else}}
	{{$authorID:=0}}{{$message:=.nil}}{{$channel:=.nil}}{{$rest:=""}}{{$command:=""}}{{$type:=""}}{{$SNum:=0}}
	{{$Syntax =print .Cmd " <Suggestion_ID> <Message/Arguments>"}}

	{{with reFindAllSubmatches (print `(?si)\A(?:` $Escaped_Prefix `\s?|\S+\s*)(?:(del|edit)\w+|\w+\s+(\w+))\s+(\d+)\s*(.*)`) .Message.Content}}
		{{$command =lower (or (index . 0 1) (index . 0 2))}}
		{{$mID:=index . 0 3}}
		{{$rest =index . 0 4}}
		{{$globalDict.Set "mID" $mID}}
		{{template "process-suggest-msg" $globalDict}}
		{{$message =$globalDict.msg}}{{$channel =$globalDict.chan}}{{$error =$globalDict.err}}{{$type =$globalDict.type}}{{$SNum =$globalDict.SNum}}{{$authorID =$globalDict.authorID}}
	{{else}}
		{{$error ="Syntaxe incorrecte : n'a pas fourni d'ID de message valide"}}
	{{end}}

	{{if and (ne $command "comment") (not $error)}}
		{{if eq $type "Implemented"}}
			{{$error =print "Ne peut pas utiliser la commande " $command " sur les suggestions implantés."}}
		{{else if and (eq $type "Approved") (eq $command "del" "edit" "approve" "approved")}}
			{{$error =print "Ne peut pas utiliser la commande " $command " sur les suggestions approuvées."}}
		{{end}}
	{{end}}

	{{if not $error}}
		{{$embed:=structToSdict (index $message.Embeds 0)}}{{range $k,$v:=$embed}}{{if eq (kindOf $v true) `struct`}}{{$embed.Set $k (structToSdict $v)}}{{end}}{{end}}{{$embed.Author.Set "Icon_URL" $embed.Author.IconURL}}
		{{if eq $command "del"}}
			{{if or (eq $authorID .User.ID) $IS_Mod}}
				{{deleteMessage $channel $message.ID 0}}
			{{else}}
				{{$error ="Vous ne pouvez supprimer que vos propres suggestions. Assurez-vous que vous avez utilisé le bon ID de suggestion."}}
			{{end}}
		{{else if eq $command "edit"}}
			{{if eq $authorID .User.ID}}
				{{if $rest}}
					{{$embed.Set "Description" $rest}}
					{{editMessage $channel $message.ID (cembed $embed)}}
				{{else}}
					{{$error ="La suggestion éditée ne peut pas être vide"}}
				{{end}}
			{{else}}
				{{$error ="Vous ne pouvez modifier que vos propres suggestions. Assurez-vous que vous avez utilisé le bon ID de suggestion."}}
			{{end}}
		{{else if $IS_Mod}}
			{{$send_chan:=$Logging_Channel}}
			{{if ne $command "comment"}}
			{{if eq $command "dupe" "markdupe"}}
				{{$command ="Dupe"}}
			{{else if eq $command "deny"}}
				{{$command ="Denied"}}
			{{else if eq $command "approve" "approved"}}
				{{$command ="acceptée"}}{{$send_chan =$Approved_Channel}}
			{{else}}
				{{$command ="implémenté"}}{{$send_chan =$Implemented_Channel}}
			{{end}}
			{{$title := (print "Suggestion #" $SNum )}}
			{{$embed.Set "Title" $title}}
			{{end}}
			{{$co := ""}}
			{{$msg := getMessage $channel $message.ID}}
			{{range $msg.Reactions}}
				{{- $co = print "\n :white_small_square: " .Count " : " "<:" .Emoji.Name ":" .Emoji.ID ">" $co}}
		 	{{- end}}
			{{$co = joinStr "" "**__Résultat du vote__** : " $co}}

			{{if eq $command "Dupe"}}
				{{$Syntax ="<Suggestion_ID> <Original_Suggestion_ID>"}}
				{{with $rest}}
					{{$globalDict.Set "mID" .}}{{$globalDict.Set "msg" $.nil}}
					{{template "process-suggest-msg" $globalDict}}
					{{if not $globalDict.err}}
						{{if lt $globalDict.SNum $SNum}}
							{{if $globalDict.msg}}
								{{$embed.Set "Description" (print $embed.Description "\n\n**Ce message a été marqué comme étant un doublon de :\n**https://discordapp.com/channels/" $.Guild.ID "/" $globalDict.chan "/" .)}}
								{{deleteMessage $channel $message.ID 0}}
								{{sendMessage $send_chan (complexMessage "content" (print "<@" $authorID "> | La suggestion suivante a été marqué comme doublon : ") "embed" $embed)}}
							{{else}}
								{{$error =print "Message ID de la suggestion original invalide : `" $rest "`"}}
							{{end}}
						{{else}}
							{{$error ="La suggestion originale doit être plus vieille que le doublon."}}
						{{end}}
					{{else}}
						{{$error =print "Suggestion originale : " $globalDict.err}}
					{{end}}
				{{else}}
					{{$error ="N'a pas fourni d'ID valide pour le message de suggestion original"}}
				{{end}}
			{{else if eq $command "Denied"}}
				{{deleteMessage $channel $message.ID 0}}
				{{sendMessage $send_chan (complexMessage "content" (joinStr "" "<@" $authorID "> | La suggestion a été refusé " $rest "\n" $co) "embed" $embed)}}
			{{else if eq $command "comment"}}
				{{template "handle-comments" (sdict "embed" $embed "comment" $rest "user" $.User)}}
				{{editMessage $channel $message.ID (cembed $embed)}}
			{{else}}
				{{template "handle-comments" (sdict "embed" $embed "comment" $rest "user" $.User)}}
				{{$embed.Footer.Set "Text" (print $command " Par : " .User.Username " - " .User.ID " ● " $embed.Footer.Text)}}
				{{deleteMessage $channel $message.ID 0}}
				{{if ne $send_chan $Logging_Channel}}{{sendMessage  $send_chan (cembed $embed)}}{{end}}
				{{sendMessage $Logging_Channel (complexMessage "content" (print print "<@" $authorID "> |La suggestion suivante a été " $command "\n" $co) "embed" $embed)}}
			{{end}}
		{{else}}
			{{$error ="Vous devez être un modérateur ou un admin pour utiliser les commandes d'administrations."}}
		{{end}}
	{{end}}
{{end}}
{{end}}

{{if not (or $Attachments $Img_Set)}}{{deleteTrigger 20}}{{end}}{{deleteResponse 5}}
{{if $error}}
	{{$ID:=sendMessageRetID nil (cembed "title" "Erreur" "color" 0xFF0000 "description" (print "**Erreur:** " $error "\n\n**Syntaxe:** `" $Syntax "`"))}}
	{{deleteMessage nil $ID 25}}
{{else}}
Done :+1:
{{end}}

{{define "handle-comments"}}
	{{if and (not .embed.Fields) .comment}}{{.embed.Set "Description" (print .embed.Description "\n\n**__Comment:__**")}}{{else if not .comment}}{{.embed.Set "Description" (reReplace  `\n\n\*\*__Comment:__\*\*\z` .embed.Description "")}}{{end}}
	{{if .comment}}{{.embed.Set "Fields" (cslice (sdict "name" (print "<@" .user.ID ">") "value" .comment))}}{{else}}{{.embed.Set "Fields" cslice}}{{end}}
{{end}}

{{define "process-suggest-msg"}}
	{{$err:=""}}
	{{range $k,$v:=.chans}}
		{{if not $.msg}}{{with getMessage $k $.mID}}{{$.Set "msg" .}}{{$.Set "chan" $k}}{{end}}{{end}}
	{{end}}

	{{with .msg}}
		{{with .Embeds}}
			{{with (index . 0).Footer}}
				{{with reFindAllSubmatches `(?s).*Author ID - (\d{17,19})\z` .Text}}
					{{$.Set "authorID" (toInt64 (index . 0 1))}}
				{{else}}
					 {{$err ="Message de suggestion invalide."}}
				{{end}}
			{{else}}
				{{$err ="Message de suggestion invalide."}}
			{{end}}
			{{with reFindAllSubmatches `\A(?:(Suggestion)|(Approved) Suggestion|(Implemented) Suggestion) #(\d+).*` (index . 0).Title}}
				{{$.Set "type"  (or (index . 0 1) (index . 0 2) (index . 0 3))}}
				{{$.Set "SNum" (toInt (index . 0 4))}}
			{{else}}
				{{$err ="Message de suggestion invalide. Erreur 2"}}
			{{end}}
		{{else}}
			{{$err ="Message de suggestion invalide."}}
		{{end}}
	{{else}}
		{{$err =print "ID incorrect : `" $.mID "`"}}
	{{end}}
	{{.Set "err" $err}}
{{end}}
{{deleteTrigger 1}}