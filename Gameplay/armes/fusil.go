{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}

{{/* Groupe dictionnaire */}}
{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Get player */}}


{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id:= .User.ID}}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}

{{$img := "https://i.imgur.com/YeIsRmw.png"}}


{{/* get PA */}}
{{$pa := $groupe.Get (str $id)}}
{{if not $pa}}
	{{$groupe.Set (str $id) 4}}
{{end}}

{{/* Dict for weapon */}}
{{$arme := sdict}}
{{with (dbGet $id "arme")}}
	{{$arme = sdict .Value}}
{{end}}

{{$desc := ""}}

{{if gt $pa 0}}
	{{if not ($arme.Get "fusil")}}
		{{$arme.Set "fusil" 1}}
		{{$y := $arme.Get "fusil"}}
		{{$x := sub 12 $y}}
		{{if lt (toFloat $y) (toFloat 12)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $x)) "/12 charges de fusil.")}}
		{{else if eq (toFloat $y) 0}}
			{{$desc = "Dernière charge utilisée."}}
		{{else}}
			{{ $desc = "Fusil vide."}}
		{{end}}
	{{else}}
		{{$arme.Set "fusil" (add ($arme.Get "fusil") 1)}}
		{{$y := $arme.Get "fusil"}}
		{{$x := toFloat (sub 12 $y)}}
		{{if lt (toFloat $y) (toFloat 12)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $x)) "/12 charges de fusil.")}}
		{{else if eq (toFloat $y) 0}}
			{{$desc = "Dernière charge utilisée."}}
		{{else}}
			{{ $desc = "Fusil vide."}}
		{{end}}
	{{end}}
{{else}}
	{{$desc = "PA insuffisants pour réaliser l'action."}}
{{end}}

{{$embed := cembed
"author" (sdict "name" $user "icon_url" $img)
"color"  0x6CAB8E
"description" $desc}}
{{ $idM := sendMessageRetID nil $embed }}
{{deleteMessage nil $idM 30}}

{{dbSet $id "arme" $arme}}