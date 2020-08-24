{{$chancommande := "701373579593252944"}}
{{$combat := 701373579593252944}}
{{$col := 16777215}}
{{$p := 0}}
{{$r := .Member.Roles}}
{{range .Guild.Roles}}
	{{if and (in $r .ID) (.Color) (lt $p .Position)}}
	{{$p = .Position}}
	{{$col = .Color}}
	{{end}}
{{end}}
{{$link := ""}}

{{$commande := ""}}

{{if .CmdArgs}}
	{{$commande = index .Args 0}}
	{{if or (eq $commande "?arme") (eq $commande "?armes")}}
		{{$message := getMessage 701377081908527124 743999494328287363 }}

		{{$embed := cembed
    		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content )
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))}}
			{{sendMessage nil $embed}}
			{{deleteTrigger 1}}

	{{else if or (eq $commande "?résumé") (eq $commande "?resume") (eq $commande "?résume") (eq $commande "?resumé") (eq $commande "?résumés") (eq $commande "?resumes") (eq $commande "?résumes") (eq $commande "?resumés")}}
		{{$embed := cembed
			"title" "Commande de base"
			"color" $col
			"description" "▫️ Pour un dé simple : `$d Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n ▫️ Pour un dé avec bonus extérieur à un implant : `$d -bonus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n ▫️ Pour un dé avec malus : `$d malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n ▫️ Pour un dé avec malus, bonus : `$d -bonus malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n:\nwhite_small_square **Autre type d'action :** Pour les utiliser, remplacer `$d` par la commande suivante : \n <:tr:724626754282717194> :small_blue_diamond: `$poison` : poison\n <:tr:724626754282717194> :small_blue_diamond: `$soin` : Soin\n <:tr:724626754282717194> :small_blue_diamond: `$malus` : Malus (de statistiques)"}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?dés") (eq $commande "?dice") (eq $commande "?dé") (eq $commande "?de") (eq $commande "?dices")}}

		{{$message := getMessage $combat 727992973131776020 }}
		{{$embed := cembed
	   	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content )
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "dégât") (eq $commande "?dégat") (eq $commande "?degât") (eq $commande "?degat") (eq $commande "?dégâts") (eq $commande "?dégats") (eq $commande "?degâts") (eq $commande "?degats")}}

		{{$message := getMessage $combat 701374206268538960}}
		{{$embed := cembed
	   	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content )
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?position") (eq $commande "?positions")}}
		{{$embed :=cembed
		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
		"color" $col
		"description" "▫️  *Abrité* : +3 en précision pour les attaquants, mais ne peut ni attaquer, ni esquiver.\n▫️ *A découvert* : -1 aux jets d'attaques (pour les attaquants).\n▫️  *Duel* : -1 aux jets d'attaques (pour les attaquants), et +2 en précision si visée sur les autres rangs."
		"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?recharge") (eq $commande "?recharger") (eq $commande "?charger") (eq $commande "?rechargé") (eq $commande "?chargé")  (eq $commande "?charge") (eq $commande "?charger")}}

		{{$message := getMessage 734838748721840188 735869281698447463}}
		{{$embed := cembed
	   	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))
			"description" (joinStr "" $message.Content )
			}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if eq $commande "?shop" "?store"}}

		{{$message := getMessage 734838748721840188 736713337571901570}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content)
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if eq $commande "?pa" "?PA"}}

		{{$message := getMessage $combat 739903306683646105}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))
			"description" (joinStr "" $message.Content )}}
			{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if eq $commande "?jdr"}}
		→ https://www.jdr-system.ovh/
		{{deleteTrigger 1}}

	{{else if eq $commande "?archive"}}
		→ https://mara-li.github.io/Archive-RP/Projet_Nucleus/readme.html
		{{deleteTrigger 1}}
	{{else if eq $commande "?wiki"}}
		→ https://github.com/Mara-Li/Projet-Nucleus-Custom-Command/wiki
		{{deleteTrigger 1}}
		{{else if eq $commande "?recensement" "?recens" "?perso" "?char" "?charlist" "?googledoc" "?doc"}}
			→ https://docs.google.com/spreadsheets/d/1k_7glSefzeAqWCFu9F3lWfYfCEw4cIq_ijWz2z-PwnU/edit?usp=sharing
			{{deleteTrigger 1}}

	{{else if eq $commande "?echange" "?give" "?échange"}}

		{{$message := getMessage 734838748721840188 739978392409079839}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String) "icon_url" (.User.AvatarURL "512"))
			"description" (joinStr "" $message.Content )}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if eq $commande "?template"}}
		→ https://docs.google.com/document/d/1CX4ye8loV4d34BOwmRlb-nOt8SMA4VvbFy4-_MHX5A8/edit?usp=sharing
		→ Fichier : Créer une copie
		→ __Remplacer l'image__ : clic droit (sur l'image) → Remplacer l'image.
		→ __ Remplir une case__ : Pot de peinture
		{{deleteTrigger 1}}

	{{else if eq $commande "?prix"}}
		{{$message := getMessage 722755391498485800 747503553307869334}}
		{{sendMessage nil (index $message.Embeds 0)}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?help") (eq $commande "?snippet") (eq $commande "?all") }}
		{{$embed := cembed
			"title" "Liste des aides disponibles"
			"description" "▫️ `?(armes|arme)`\n▫️ `?résumé`\n▫️ `?dés`\n▫️ `?position`\n▫️ `?(dégâts|dégât|dégat|dégats)`\n▫️ `?charge`\n▫️ `?shop`\n▫️`?(pa|PA)`\n▫️ `?archive`\n▫️`?(shop|store)`\n ▫️ `?template`\n▫️ `?(recen(sement?)|char(list?)|googledoc|doc)`\n▫️ `?((e|é)change|give)`\n ▫️ `?prix`\n ▫️ `?wiki`\n\n :white_medium_square: **Pour afficher cette liste** : `?(all|snippet|help)`\n\n [Wiki des commandes](https://github.com/Mara-Li/Projet-Nucleus-Custom-Command/wiki)"}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}
	{{end}}
{{end}}
