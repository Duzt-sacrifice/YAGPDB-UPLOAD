{{/* Trigger type: command  Trigger: moderate */}}

{{/* Variables */}}
{{$mod := ROLE-ID}}
{{$admin := ROLE-ID}}

{{/* Don't Mess with this unless you know what you are doing */}} {{/* Command */}}

{{if or (hasRoleID $admin) (hasRoleID $mod)}}
{{deleteTrigger 10}}
    {{if eq (len .CmdArgs) 1}}
        {{$user := (userArg (index .CmdArgs 0))}}
        {{$embed := cembed "title" "Moderation Menu" "description" (joinStr "" "❗ - Warn\n🔇 - Mute\n🔈 - Unmute\n👢 - Kick\n🔨 - Ban\n❌ - Exit\n \\> React to choose what action you want to perform < ") "color" 77}}
        {{$a := sendMessageRetID nil $embed}}
        {{dbSet .User.ID "modmenu" (joinStr ":" (toString $a) (toString $user.ID))}}
        {{addMessageReactions nil $a "❗" "🔇" "🔈" "👢" "🔨" "❌"}}
    {{end}}
{{end}}