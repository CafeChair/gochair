{{define "navbar"}}
		<a class="navbar-brand" href="/">GoChair</a>
		<div>
			<ul class="nav navbar-nav">
				<li {{if .IsIndex}}class="active"{{end}}><a href="/">DashBoard</a></li>
				<li {{if .IsQdns}}class="active"{{end}}><a href="/qdns">HttpDNS</a></li>
				<li {{if .IsRunCmd}}class="active"{{end}}><a href="/runcmd">RunCmd</a></li>
				<li {{if .IsCron}}class="active"{{end}}><a href="/cron">Crontab</a></li>
				<li {{if .IsRunJob}}class="active"{{end}}><a href="/runjob">RunJob</a></li>
				<li {{if .IsAgent}}class="active"{{end}}><a href="/agent">Agent</a></li>
				<li {{if .IsTodo}} class="active" {{end}}><a href="/todo">Todo</a></li>
			</ul>
		</div>
{{end}}