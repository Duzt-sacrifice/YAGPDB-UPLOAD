{{/* Trigger Type: command  Trigger: login

	This command is a login command for entering a password and recieving a role based on it for a specified amount of time. Theres an optional max time setting (basically anything you see in green below this message)  */}}

{{$args := parseArgs 2 "Invalid Arguments Provided, Command usage: `-login password time`"
    (carg "string" "Password Used")
    (carg "string" "Time in area")}}

{{/* $maxtime := seconds */}}
{{$pwu := ($args.Get 0)}}
{{$time := ($args.Get 1)}}
{{$passwords := cslice "password1" "password2" "password3"}}
{{$roles := sdict "password1" role-to-be-added "password2" role-to-be-added "password3" role-to-be-added}}

{{if (inFold $passwords $pwu)}}
    {{$duration := (toDuration $time).Seconds}}
    {{/* if not (gt (toInt $duration) (toInt $maxtime)) */}}
        {{addRoleID ($roles.Get $pwu)}}
        {{removeRoleID ($roles.Get $pwu) $duration}}
        {{sendMessage nil (joinStr "" "Access Granted: " .User.Username " - " $pwu ": " $time ".")}}
    {{/* else */}}
        {{/* Please dont specify more than 12 hours, Command usage: `-login password time` */}}
    {{/* end */}}
{{else}}
    Please choose a valid password, Command usage: `-login password time`
{{end}}