{{$n:="**Biocomposant :** Un liquide jaunâtre se trouvant dans un tube en plastique. Utilisé pour fabriquer tout type d'analgésique dans les fabricateurs biologiques."}}{{$o:="**Cellule bionotropique** : Un liquide verdâtre, assez rare, qui permet la fabrication d'implant bionotropique."}}{{$p:="**Liquide cytomorphe** : Une sorte de liquide grisâtre se trouvant dans un tube en plastique, qui permet de fabriquer des modules dans les fabricateurs."}}{{$q:="**Substrat ferreux** : Un liquide rouge que l'on utilise souvent pour fabriquer des armes. Très rares et couteux."}}{{$r:="**Composant universel** : Un cube semi-solide de couleur bleu. On en trouve quasiment jamais, et ils se vendent une fortune au marchée noir. Il permet de créer n'importe quel objet dans n'importe quel fabricateur : module, implant, arme, analgésique..."}}{{$s:="**Rixolam** : Un comprimé en gélule blanc et vert, restaurant 25 PV après ingestion."}}{{$g:="**Bandage** : Des bandes de tissus assez simples et vieilles comme Ophéris, qui permette de restaurer 10 pv et réduire de deux tours un malus."}}{{$t:="**Eufyrusant** : Un boost se présentant comme un comprimé rose, destiné aux PSI,  qui permet d'augmenter les capacités de 8% durant trois tours."}}{{$u:="**Implant temporaire** : Un patch triangulaire transparent qui se place sur la peau, agissant sur une seule caractéristique durant trois tours. La caractéristique est décidé lors de sa création."}}{{$w:="**Soma** : Un comprimé rouge, permettant d'annuler tous les cooldown d'un module ou d'un PSI."}}{{$x:="**Injection Eufyrusant** : Fabriqué par la même corporation que le comprimé, cette injection est destinée aux personnes ayant un module, et qui augmente les capacités d'un module de 8% pendant 3 tours."}}{{$y:="**Sirop de betapropyl** : Un sirop hyper sucrée et légèrement addictif, qui annule tous les malus dont la personne était victime."}}{{$z:="**Xenox** : Un comprimé vert très précieux qui permet d'avoir, pendant trois tours, l'effets d'un implant sur trois caractéristiques au choix."}}{{$A:="**Grenade nécrotique** : Une grenade qui provoque des nécroses sur la peau de la personne l'ayant reçue. A manipuler avec précautions."}}{{$B:="**Liquide antifongique** : Il fait fondre les cellules épaisses de protection d'un Mush."}}{{$C:="**Gaz anesthésiant de combat : ** Bonne nuit !"}}{{$h:="**Sang Etherique** : Un petit nom donnée au sang particulier d'Ether. Il peut être utilisé dans de nombreux objets différents. Dans tous les cas : il est toxique."}}{{$D:="**Huile carotoxinique :** Une huile issue des toxine d'une araignée, destinée aux armes blanches. Une fois dans les blessures, elle augmente la douleur."}}{{$E:="**Huile digestive** : Une arme très redoutable, composée d'une protéase qui digère les armures."}}{{$G:=""}}{{$H:="722755391498485800"}}{{$I:=""}}{{$J:=""}}{{$F:=""}}{{$K:=print "Recherche faite par " (getMember .User.ID).Nick}}{{$L:=true}}{{$M:="https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png"}}{{$N:="https://i.imgur.com/o557fMx.png"}}{{if .CmdArgs}}{{$O:=lower (index .CmdArgs 0)}}{{if eq $O "cb"}}{{$G =$o}}{{else if eq $O  "bc"}}{{$G =$n}}{{else if eq $O "lc"}}{{$G =$p}}{{else if eq $O "sf"}}{{$G =$q}}{{else if eq $O "cu"}}{{$G =$r}}{{else if eq $O "rixolam"}}{{$G =$s}}{{else if eq $O "bandage" "bandages"}}{{$G =$g}}{{else if eq $O "eufyrusant" "eufy"}}{{$G =$t}}{{$G =$x}}{{else if eq $O "implants temporaires" "implant temporaire"}}{{$G =$u}}{{else if eq $O "soma"}}{{$G =$w}}{{else if eq $O "sirop" "betapropyl" "sirop de betapropyl"}}{{$G =$y}}{{else if eq $O "xenox"}}{{$G =$z}}{{else if eq $O "grenade" "grenade nécrotique" "grenade necrotique" "nécrotique" "necrotique" "necro" "nécro"}}{{$G =$A}}{{else if eq $O "liquide antifongique" "liquide" "antifongique"}}{{$G =$B}}{{else if eq $O "anesthésiant" "gaz anesthésiant" "gaz"}}{{$G =$C}}{{else if eq $O "sang etherique" "sang ethérique" "sang" "éthérique" "etherique" "étherique" "ethérique" "sang étherique"}}{{$G =$h}}{{else if eq $O "carotoxinique" "caro" "huile carotoxinique"}}{{$G =$D}}{{else if eq $O "huile digestive" "digestive"}}{{$G =$E}}{{else if eq $O "huile" "huiles"}}{{$G =joinStr "" ":white_small_square:" $E "\n :white_small_square:" $D}}{{else if eq $O "inventaire"}}{{$F ="736003604221001790"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "composant"}}{{$F ="736003719631601775"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J =(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "analgésique"}}{{$F ="736003785398288404"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "arme biologique" "armes biologiques" "armes biologique" "arme biologiques"}}{{$F ="736004347556528182"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "armes" "arme"}}{{$F ="736004668697870336"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "module" "modules"}}{{$F ="736004830883086457"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "implants" "implant"}}{{$F ="736005099360485388"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "charge" "charges"}}{{$F ="736006005355315200"}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else if eq $O "faible"}}{{if eq (len .CmdArgs) 2}}{{$e:=reFind `(?i)(analgésiques?|(Armes? biologiques?)|(armes? bio)|Armes?)` (index .CmdArgs 1)}}{{$e =lower $e}}{{if eq $e "analgésique" "analgésiques"}}{{$F ="736005431582916638"}}{{else if eq $e "arme biologique" "armes biologiques" "arme biologiques" "armes biologique" "arme bio" "armes bio"}}{{$F ="736005737452404778"}}{{else if eq $e "arme" "armes"}}{{$F ="736006328668913684"}}{{end}}{{$j:=getMessage $H $F}}{{$I =(index $j.Embeds 0).Description}}{{$J:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $H "/" $F ")")}}{{$G =(joinStr "" $I )}}{{else}}{{$L =true}}
{{$G ="Vous cherchez un objet à bas coût ? Mais lequel ?\n\n Pour rappel il existe :\n :white_small_square: Les armes (`?search faible arme(s)`)\n :white_small_square: Les armes biologiques (`?search faible \"arme bio\"`) \n :white_small_square: Les analgésiques (`?search faible analgésique`)"}}
{{end}}{{$L =true}}{{else}}{{$L =false}}{{$G ="Cet élément ne figure pas dans ma base de données."}}{{end}}{{else}}{{$L =false}}{{$G ="Pourquoi vous ouvrez l'encyclopédie si ce n'est pas pour y chercher quelque chose ? "}}{{end}}

{{if eq $L true}}
{{$c:=cembed
"author" (sdict "name" "[Sola-UI] Base de donnée : Objet" "icon_url" $M)
"thumbnail" (sdict "url" $N)
"description" $G
"footer" (sdict "text" $K )
"color" 0x94CAF0}}
{{sendMessage nil $c}}

{{else if eq $L false}}
{{$c:=cembed
"author" (sdict "name" "[Sola-UI] Base de donnée : Objet" "icon_url" $M)
"thumbnail" (sdict "url" $N)
"description" (joinStr " " $G "\n\n Voici le sommaire : \n▫️ **Inventaire** : `?search inventaire` \n▫️ **Composant** : `?search composant`\n▫️ **Analgésique** : `?search analgésique`\n▫️ **Armes biologiques** : `?search \"Arme biologique\"`\n▫️ **Armes** : `?search arme`\n▫️ **Modules** : `?search module`\n▫️ **Implants** : `?search implants`\n▫️ **Charges** : `?search charge` \n ▫️ **Objet à bas coût** : Voir à la page `?search faible ` pour plus d'informations ")
"fields" (cslice
(sdict "name" "_ _" "value" "_ _" "inline" false)
(sdict "name" "Composants" "value" "▫️ **Biocomposant** : `?search bc`\n▫️ **Cellule bionotropique** : `?search cb`\n▫️ **Cellule cytomorphe** : `?search lc`\n▫️ **Substrat ferreux** : `?search sf`\n▫️ **Composant universel** : `?search cu`" "inline" false)
(sdict "name" "_ _" "value" "_ _" "inline" false)
(sdict "name" "Analgésique" "value" ":white_small_square:**Rixolam :** `?search rixolam`\n:white_small_square:**Bandage :** `?search bandage(s)`\n:white_small_square:**Eufyrusant :**`?search eufy(rusant)`\n:white_small_square:**Implant :** `?search \"Implant(s) temporaire(s)\"`\n:white_small_square:**Soma :** `?search soma`\n:white_small_square:**Injection Eufyrusant :** `?search Eufy(rusant)`\n:white_small_square: **Sirop de betapropyl :** `?search (sirop|betapropyl|\"sirop de betapropyl\")`\n:white_small_square: **Xenox :** `?search xenox`" "inline" false)
(sdict "name" "_ _" "value" "_ _" "inline" false)
(sdict "name" "Armes biologiques" "value" ":white_small_square: **Grenade Nécrotique :** `?search (grenade|grenade nécrotique|nécro)`\n:white_small_square: **Liquide antifongique :**`?search (antifongique|liquide|\"liquide antifongique\")`\n:white_small_square: **Gaz anesthésiant de combat :** `?search (anesthésiant|gaz|\"gaz anesthésiant\")`\n:white_small_square: **Sang Etherique :** `?search (\"sang éthérique\"|sang|éthérique)`\n:white_small_square: **Huile carotoxinique :** `?search (huile|carotonixique|caro|\"huile carotoxinique\")`\n:white_small_square: **Huile digestive :**  `?search (huile|digestive|\"huile digestive\"`" "inline" false)
)
"footer" (sdict "text" $K )
"color" 0xA75454}}
{{sendMessage nil $c}}
{{end}}
