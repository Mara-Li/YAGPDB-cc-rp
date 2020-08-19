{{/* Databases */}}
{{$tr := .User.ID}}
{{ with and (eq .ReactionMessage.Author.ID 204255221017214977) .ReactionMessage.Embeds }} {{/* Checks for validity */}}
	 {{ $embed := structToSdict (index . 0) }}
	 {{ range $k, $v := $embed }}
		 {{- if eq (kindOf $v true) "struct" }}
		 {{- $embed.Set $k (structToSdict $v) }}
		 {{- end -}}
		{{ end }}

{{$del := false}}
{{$check := ""}}
{{ $action := $.Reaction.Emoji.Name }} {{/* The action being ran */}}
{{ $validEmojis := cslice "▶️" "◀️" "🗑️" }} {{/* Valid emojis */}}
{{ $isValid := false }} {{/* Whether this is actually a valid embed / leaderboard embed */}}
{{ $page := 0 }} {{/* The current page */}}
{{if $embed.Author.Name}}
	{{$check = reFind `Inventaire` $embed.Author.Name}}
{{end}}
{{ if and (eq $check "Inventaire") $embed.Footer}} {{/* More checks */}}
	{{ $page = toInt (reFind `\d+` $embed.Footer.Text) }} {{/* We presume that this is valid, and get the page num */}}
	{{ $isValid = true }} {{/* Yay, it is valid */}}
{{else if and (eq $check "Inventaire")}}
	{{$isValid = true}}
	{{$page = 1}}
{{ end }}

{{ if and (in $validEmojis $action) $isValid $page }}
{{$id := reFind `(\#\S*)` $embed.Footer.Text}}
{{$id = (toInt (joinStr "" (split $id "#")))}}
{{$user := ""}}
{{if eq (toInt $id) $tr}}
	{{$user = (getMember $id).Nick}}
{{else}}
	{{$user = (dbGet (toInt $id) "rerollName").Value}}
{{end}}



{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$serverEco := sdict}}
{{with (dbGet $.Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}


{{$desc := "Ton inventaire est vide ! Si le shop est ouvert, tu peux aller acheter des trucs !"}}
{{$footer := print "Page: 1 / 1 | #" $id }}
{{$cslice := cslice}}
{{range $k,$v := $inv}}
	{{$cslice = $cslice.Append (printf " :white_small_square: ** %-10v** : [%v]" $k $v)}}
{{end}}
{{$author := (joinStr " " "Inventaire de :" (title $user))}}

{{/* Pagination */}}
		{{ deleteMessageReaction nil $.ReactionMessage.ID $.User.ID $action }}
	{{ if eq $action "▶️" }}
		{{ $page = add $page 1 }} {{/* Update page according to emoji */}}
	{{ else if eq $action "◀️"}}
		{{ $page = sub $page 1 }}
		{{if le $page 1}}
			{{$page =1}}
		{{end}}
	{{else}}
		{{$del = true}}
		{{$page = 1}}
		{{deleteMessage nil $.ReactionMessage.ID 1}}
	{{ end }}

	{{$start := ""}}
	{{$stop := ""}}
	{{$end := ""}}
	{{if $cslice}}
		{{$start = (mult 10 (sub $page 1))}}
		{{$stop = (mult $page 10)}}
		{{$end = roundCeil (div (toFloat (len $cslice)) 10)}}
		{{$data := ""}}
		{{if ge $stop (len $cslice)}}
			{{$stop = (len $cslice)}}
		{{end}}
		{{if ne $page 0}}
			{{if and (le $start $stop) (ge (len $cslice) $start) (le $stop (len $cslice))}}
				{{range (seq $start $stop)}}
				{{$data = (print $data "\n" (index $cslice .))}}
				{{end}}
{{$footer = print "Page: " $page " / " $end " | #" $id }}
			{{else}}
				{{$data = "Il n'y a rien ici..."}}
{{$footer = print "Page: " $page " / " $end " | #" $id }}
			{{end}}
		{{else}}
			{{$data = "Il n'y a rien ici..."}}
{{$footer = print "Page: " $page " / " $end " | #" $id }}
		{{end}}
			{{$desc = print "" $data ""}}
	{{end}}


	{{if eq $del false}}
		{{editMessage nil $.ReactionMessage.ID (cembed "author" (sdict "name" $author "icon_url" "https://i.imgur.com/iUmz9Gi.png") "color" 0x8CBAEF "description" $desc "footer" (sdict "text" $footer) )}}
	{{else}}
		{{deleteMessage nil $.ReactionMessage.ID 1}}
		{{dbDel $id "rerollName"}}
	{{end}}
{{end}}
{{end}}