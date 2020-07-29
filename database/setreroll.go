{{if .CmdArgs}}
	{{$name := toRune (lower (index .CmdArgs 0))}}
	{{$user :=""}}
	{{range $name}}
	{{- $user = (print $user .)}}
	{{- end}}
	{{$user = (toInt $user)}}

	{{$stats := sdict}}
	{{with (dbGet $user "stats")}}
		{{$stats = sdict .Value}}
	{{end}}

	{{if eq (index .CmdArgs 1) "-all"}}
		{{$stats.Set "force" (toInt (index .CmdArgs 2)) }}
		{{$stats.Set "i_force" (toInt (index .CmdArgs 3)) }}
		{{$stats.Set "endurance" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "i_endu" (toInt (index .CmdArgs 5))}}
		{{$stats.Set "agi" (toInt (index .CmdArgs 6))}}
		{{$stats.Set "i_agi" (toInt (index .CmdArgs 7))}}
		{{$stats.Set "preci" (toInt (index .CmdArgs 8))}}
		{{$stats.Set "i_preci" (toInt (index .CmdArgs 9))}}
		{{$stats.Set "intelligence" (toInt (index .CmdArgs 10))}}
		{{$stats.Set "i_intel" (toInt (index .CmdArgs 11))}}
		{{$stats.Set "karma" (toInt (index .CmdArgs 12))}}
		{{dbSet $user "stats" $stats}}

**Statistiques de {{index .CmdArgs 0}}**
	:white_small_square: Force : {{($stats.Get "force")}}
	:white_small_square: Endurance : {{($stats.Get "endurance")}}
	:white_small_square: Agilité : {{($stats.Get "agi")}}
	:white_small_square: Précision : {{($stats.Get "preci")}}
	:white_small_square: Intelligence : {{($stats.Get "intelligence")}}
	:white_small_square: Karma : {{($stats.Get "karma")}}

**Implant de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{($stats.Get "i_force")}}
	:white_small_square: Endurance : {{($stats.Get "i_endu")}}
	:white_small_square: Agilité : {{($stats.Get "i_agi")}}
	:white_small_square: Précision : {{($stats.Get "i_preci")}}
	:white_small_square: Intelligence : {{($stats.Get "i_intel")}}

	{{else if eq (index .CmdArgs 1) "-stats"}}
		{{$stats.Set "force" (toInt (index .CmdArgs 2)) }}
		{{$stats.Set "endurance" (toInt (index .CmdArgs 3))}}
		{{$stats.Set "agi" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "preci" (toInt (index .CmdArgs 5))}}
		{{$stats.Set "intelligence" (toInt (index .CmdArgs 6))}}
		{{$stats.Set "karma" (toInt (index .CmdArgs 7))}}
		{{dbSet $user "stats" $stats}}

**Statistiques de {{index .CmdArgs 0}}**
	:white_small_square: Force : {{($stats.Get "force")}}
	:white_small_square: Endurance : {{($stats.Get "endurance")}}
	:white_small_square: Agilité : {{($stats.Get "agi")}}
	:white_small_square: Précision : {{($stats.Get "preci")}}
	:white_small_square: Intelligence : {{($stats.Get "intelligence")}}
	:white_small_square: Karma : {{($stats.Get "karma")}}

	{{else if eq (index .CmdArgs 1) "-implant"}}
		{{$stats.Set "i_force" (toInt (index .CmdArgs 2)) }}
		{{$stats.Set "i_endu" (toInt (index .CmdArgs 3))}}
		{{$stats.Set "i_agi" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "i_preci" (toInt (index .CmdArgs 5))}}
		{{$stats.Set "i_intel" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "stats" $stats}}

**Implant de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{($stats.Get "i_force")}}
	:white_small_square: Endurance : {{($stats.Get "i_endu")}}
	:white_small_square: Agilité : {{($stats.Get "i_agi")}}
	:white_small_square: Précision : {{($stats.Get "i_preci")}}
	:white_small_square: Intelligence : {{($stats.Get "i_intel")}}
	{{else}}
		**Usage** : `$setreroll -(all|stats|implant) (nomperso) Force Endurance Agilité Précision Intelligence Karma`
	{{end}}
{{else}}
**Usage** : `$setreroll -(all|stats|implant) (nomperso) Force Endurance Agilité Précision Intelligence Karma`
{{end}}
{{deleteTrigger 1}}