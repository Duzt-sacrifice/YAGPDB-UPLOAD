{{/* Trigger Type: Regex  Trigger: ^-(dbcommands|dball|dbdel|dbcount|dbsetadd|listalldb|cleardb) 

	This custom command is a pack of database commands, to find all the commands create the custom command and run -dbcommands */}}

{{$cmd := (index .Args 0)}}

{{if eq $cmd "-dbcommands"}}
{{sendMessage nil (cembed "title" "Database Related Commands" "description" (joinStr "" "\n\n`-dbsetadd <ID> <Key> <Value>`\n\n`-dbcount`\n\n`-listalldb`\n\n`-dbdel <ID> <Key>`") "color" 77)}}
{{end}}

{{if eq $cmd "-dbdel"}}
{{/*dbdel*/}}
{{$args := parseArgs 2 "-dbdel <ID> <Key>"
	(carg "userid" "ID")
	(carg "string" "key")}}
{{dbDel ($args.Get 0) ($args.Get 1)}}
Done
{{end}}

{{if eq $cmd "-dbcount"}}
{{dbCount}}
{{end}}

{{if eq $cmd "-dbsetadd"}}
{{$args := parseArgs 3 "<ID> <Key> <Value>"
	(carg "userid" "ID")
	(carg "string" "key")
	(carg "string" "value")}}
{{dbSet ($args.Get 0) ($args.Get 1) ($args.Get 2)}}
Done
{{end}}

{{if eq $cmd "-listalldb"}}
{{$te := dbTopEntries "%" 100 0}}
{{range $te}} 
{{sendMessage nil (cembed "title" "Database Entry" "description" (joinStr "" "`ID:` " .UserID " `Key:` " .Key " `Value:` " .Value) "color" 77 )}}
{{end}}
Done
{{end}}
{{deleteTrigger 5}} 