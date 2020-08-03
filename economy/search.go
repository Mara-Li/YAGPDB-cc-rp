
{{$bc:="**Biocomposant :** Un liquide jaunâtre se trouvant dans un tube en plastique. Utilisé pour fabriquer tout type d'analgésique dans les fabricateurs biologiques."}}
{{$cb:="**Cellule bionotropique** : Un liquide verdâtre, assez rare, qui permet la fabrication d'implant bionotropique."}}
{{$lc:="**Liquide cytomorphe** : Une sorte de liquide grisâtre se trouvant dans un tube en plastique, qui permet de fabriquer des modules dans les fabricateurs."}}
{{$sf:="**Substrat ferreux** : Un liquide rouge que l'on utilise souvent pour fabriquer des armes. Très rares et couteux."}}
{{$cu:="**Composant universel** : Un cube semi-solide de couleur bleu. On en trouve quasiment jamais, et ils se vendent une fortune au marchée noir. Il permet de créer n'importe quel objet dans n'importe quel fabricateur : module, implant, arme, analgésique..."}}



{{$ri:="**Rixolam** : Un comprimé en gélule blanc et vert, restaurant 25 PV après ingestion." }}

{{$ban:="**Bandage** : Des bandes de tissus assez simples et vieilles comme Ophéris, qui permette de restaurer 10 pv et réduire de deux tours un malus."}}

{{$eu:="**Eufyrusant** : Un boost se présentant comme un comprimé rose, ou sous forme d'injection, qui permet d'augmenter les capacités de 8% durant trois tours."}}

{{$so:="**Soma** : Un comprimé rouge, permettant d'annuler tous les cooldown d'un module ou d'un PSI."}}

{{$ro:="**Sirop de betapropyl** : Un sirop hyper sucrée et légèrement addictif, qui annule tous les malus dont la personne était victime."}}

{{$xe:="**Xenox** : Un comprimé vert très précieux qui permet d'avoir, pendant trois tours, l'effets d'un implant sur trois caractéristiques au choix." }}


{{$gr:="**Grenade nécrotique** : Une grande qui provoque des nécroses sur la peau de la personne l'ayant reçue. A manipuler avec précautions. (10% bonus projectile)"}}

{{$fn:="**Liquide antifongique** : Il fait fondre les cellules épaisses de protection d'un Mush. (Annule un bouclier mush pendant 3 tours)"}}

{{$th:="**Gaz anesthésiant de combat : ** Endort la cible. "}}

{{$sng:="**Sang Etherique** : Un petit nom donnée au sang particulier d'Ether. Il peut être utilisé dans de nombreux objets différents. Dans tous les cas : il est toxique. (15% bonus)"}}

{{$cr:="**Huile carotoxinique :** Une huile issue des toxine d'une araignée, destinée aux armes blanches. Une fois dans les blessures, elle augmente la douleur. (Malus de caractéristiques sur tout le combat)"}}

{{$di:="**Huile digestive** : Une arme très redoutable, composée d'une protéase qui digère les armures. (Annule l'armure sur tout le combat)"}}

{{$cr := "**Charge creuse** : Charge ayant l'effet d'un burst, avec les mêmes limitations."}}
{{$cp := "**Charge perforante** : Charge ayant l'effet d'une attaque perforante, avec les mêmes limitations."}}
{{$ce := "**Charge explosive** : Inflige des dégâts normaux sur tout le rang."}}
{{$cd := "**Charge dégénérescente** : Charge ayant l'effet d'une altération de statistiques."}}
{{$cg := "**Charge gangrenante** : Charge ayant l'effet d'un poison, avec les mêmes limitations."}}

{{$d:=""}}
{{$v:="722755391498485800"}}
{{$m:=""}}
{{$k:=""}}
{{$id:=""}}
{{$f:=print "Recherche faite par " (getMember .User.ID).Nick }}
{{$b:=true}}
{{$a:="https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png"}}
{{$l:="https://i.imgur.com/o557fMx.png"}}
{{$fu:= true}}

{{if .CmdArgs}}
	{{$i:=lower (index .CmdArgs 0)}}
	{{if eq $i "cb"}}
	  {{$d =$cb}}
	{{else if eq $i "faible"}}
		{{$fu = false}}
	{{else if eq $i  "bc"}}
	  {{$d =$bc}}
	{{else if eq $i "lc"}}
	  {{$d =$lc}}
	{{else if eq $i "sf"}}
	  {{$d =$sf}}
	{{else if eq $i "cu"}}
	  {{$d =$cu}}
	{{else if eq $i "rixolam"}}
		{{$d =$ri}}
	{{else if eq $i "bandage" "bandages"}}
		{{$d =$ban}}
	{{else if eq $i "eufyrusant" "eufy"}}
		{{ $d =$eu}}
	{{else if eq $i "soma"}}
		{{$d =$so}}
	{{else if eq $i "sirop" "betapropyl" "sirop de betapropyl"}}
		{{$d =$ro}}
	{{else if eq $i "xenox"}}
		{{$d =$xe}}
	{{else if eq $i "grenade" "grenade nécrotique" "nécrotique" "nécro"}}
		{{$d =$gr}}
	{{else if eq $i "liquide antifongique" "liquide" "antifongique"}}
		{{$d =$fn}}
	{{else if eq $i "anesthésiant" "gaz anesthésiant" "gaz"}}
		{{$d =$th}}
	{{else if eq $i "sang éthérique" "sang" "éthérique"}}
		{{$d =$sng}}
	{{else if eq $i "carotoxinique" "caro" "huile carotoxinique"}}
		{{$d =$cr}}
	{{else if eq $i "huile digestive" "digestive"}}
		{{$d =$di}}
	{{else if eq $i "huile"}}
		{{$d =joinStr "" $di "\n" $cr}}
	{{else if eq $i "inventaire"}}
		{{$id ="736003604221001790"}}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "composant"}}
		{{$id ="736003719631601775"}}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k =(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "analgésique"}}
		{{$id ="736003785398288404"}}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "arme biologique" "armes biologiques" "armes biologique" "arme biologiques"}}
		{{$id ="736004347556528182" }}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "armes" "arme"}}
		{{$id ="736004668697870336" }}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "module" "modules"}}
		{{$id ="736004830883086457" }}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "implants" "implant"}}
		{{$id ="736005099360485388" }}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "charge" "charges"}}
		{{$id ="736006005355315200"}}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{else if eq $i "charge creuse" "creuse"}}
		{{$d =$cr}}
	{{else if eq $i "charge perforante" "perforante"}}
		{{$d =$cp}}
	{{else if eq $i "charge dégénérescente" "charge poison" "dégénérescente" "poison" }}
		{{$d =$cd}}
	{{else if eq $i "charge gangrénante" "gangrénante"}}
		{{$d =$cg}}

	{{else}}
		{{$b =false}}
		{{$d ="Cet élément ne figure pas dans ma base de donnée."}}
	{{end}}
{{else}}
	{{$b =false}}
	{{$d ="Pourquoi vous ouvrez l'encyclopédie si ce n'est pas pour y chercher quelque chose ? "}}
{{end}}

{{if and (eq $b true) (eq $fu true)}}
	{{$embed:=cembed
			"author" (sdict "name" "[Sola-UI] BDD : Objet haut de gamme" "icon_url" $a)
			"thumbnail" (sdict "url" $l)
			"description" $d
			"footer" (sdict "text" $f )
			"color" 0x94CAF0}}
	{{sendMessage nil $embed}}
{{else if and (eq $b false) (eq $fu true)}}
	{{$embed:=cembed
		"author" (sdict "name" "[Sola-UI] BDD : Objet | ERREUR" "icon_url" $a)
		"thumbnail" (sdict "url" $l)
		"description" (joinStr " " $d "\n\n Voici les commandes : \n▫️ **Inventaire** : `?search inventaire` \n▫️ **Composant** : `?search composant`\n▫️ **Analgésique** : `?search analgésique`\n▫️ **Armes biologiques** : `?search \"Arme biologique\"`\n▫️ **Armes** : `?search arme`\n▫️ **Modules** : `?search module`\n▫️ **Implants** : `?search implants`\n▫️ **Charges** : `?search charge`")
		"fields" (cslice
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Composants" "value" "▫️ **Biocomposant** : `?search bc`\n▫️ **Cellule bionotropique** : `?search cb`\n▫️ **Cellule cytomorphe** : `?search lc`\n▫️ **Substrat ferreux** : `?search sf`\n▫️ **Composant universel** : `?search cu`" "inline" false)
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Analgésique" "value" "▫️**Rixolam :** `?search rixolam`\n▫️**Bandage :** `?search bandage(s)`\n▫️**Eufyrusant :**`?search eufy(rusant)`\n▫️**Soma :** `?search soma`\n▫️ **Sirop de betapropyl :** `?search (sirop|betapropyl|\"sirop de betapropyl\")`\n▫️ **Xenox :** `?search xenox`" "inline" false)
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Armes biologiques" "value" "▫️ **Grenade Nécrotique :** `?search (grenade|grenade nécrotique|nécro)`\n▫️ **Liquide antifongique :**`?search (antifongique|liquide|\"liquide antifongique\")`\n▫️ **Gaz anesthésiant de combat :** `?search (anesthésiant|gaz|\"gaz anesthésiant\")`\n▫️ **Sang Etherique :** `?search (\"sang éthérique\"|sang|éthérique)`\n▫️ **Huile carotoxinique :** `?search (huile|carotonixique|caro|\"huile carotoxinique\")`\n▫️ **Huile digestive :**  `?search (huile|digestive|huile digestive`" "inline" false)
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Charges" "value" "▫️**Charge creuse** : `?search (\"charge creuse\"|creuse)`\n▫️**Charge perforante** : `?search (\"charge perforante\"|perforante)`\n▫️**Charge dégénérescente** : `?search (\"charge dégénérescente\"|charge poison|dégénérescente|poison)`\n▫️**Charge gangrénante** : `?search (\"charge gangrénante\"|gangrénante)`" "inline" false))
		"footer" (sdict "text" $f )
		"color" 0xA75454}}
		{{sendMessage nil $embed}}
{{end}}