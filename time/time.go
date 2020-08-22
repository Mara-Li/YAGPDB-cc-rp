{{/*All message database counter*/}}
{{$time := sdict }}
{{with (dbGet 0 "time")}}
	{{$time = sdict .Value}}
{{end}}

{{$number:= ($time.Get "cycle")}}

{{$msgc := (toFloat ($time.Get "mgsc"))}}
{{$pourcent := div $msgc 100}}
{{if lt $pourcent (toFloat 1)}}
	{{$pourcent = (toFloat 1)}}
{{end}}
{{ $msg := (json (div ($time.Get "message") $pourcent)) }}


{{$y := (toFloat ($time.Get "cycle"))}}
{{$txt := ($time.Get "message") }}
{{$val := (joinStr " " (toString (toInt $txt)) "message(s) dans le cycle")}} {{/* Footer message */}}
{{$day := (toString (toInt ($time.Get "jour")))}}

{{/* Thumbnail */}}
{{$nuit:="https://i.imgur.com/e04keB7.png"}}
{{$matin:="https://i.imgur.com/ZB5yT5s.png"}}
{{$midi:="https://i.imgur.com/AFOj90o.png"}}
{{$soir:="https://i.imgur.com/xSDYgqD.png"}}

{{if eq $y (toFloat 1) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0x1B3175
 		"thumbnail" (sdict "url" $nuit)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 2) }}
		{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0xDD99DF
 		"thumbnail" (sdict "url" $matin)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 3) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0xF0B535
 		"thumbnail" (sdict "url" $midi)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{/* If you want add cycle, add it here ! */}}

{{else}}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0x9593E8
 		"thumbnail" (sdict "url" $soir)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}
{{end}}

{{deleteTrigger 0}}
