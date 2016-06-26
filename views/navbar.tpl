{{define "navbar"}}
		<a class="navbar-brand" href="/">GoChair</a>
		<div>
			<ul class="nav navbar-nav">
				<li {{if .IsIndex}} class="active" {{end}}><a href="/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a></li>
				<li {{if .IsQdns}} class="active" {{end}}><a href="/qdns"><i class="fa fa-plane fa-fw"></i> HttpDNS</a></li>
				<li {{if .IsRunCmd}} class="active" {{end}}><a href="/runcmd"><i class="fa fa-plane fa-fw"></i> RunCMD</a></li>
				<li {{if .IsDatabase}} class="active" {{end}}><a href="/database"><i class="fa fa-database fa-fw"></i> DataBase</a></li>
				<li {{if .IsCron}} class="active" {{end}}><a href="/cron"><i class="fa fa-clock-o fa-fw"></i> Crontab</a></li>
				<li {{if .IsJobs}} class="active" {{end}}><a href="/jobs"><i class="fa fa-plane fa-fw"></i> Jobs</a></li>
				<li {{if .IsAgent}} class="active" {{end}}><a href="/agent"><i class="fa fa-child fa-fw"></i> Agent</a></li>
				<li {{if .IsTodo}} class="active" {{end}}><a href="/todo"><i class="fa fa-calendar fa-fw"></i> Todo</a></li>
				<li {{if .IsDoc}} class="active" {{end}}><a href="/doc"><i class="fa fa-book fa-fw"></i> Doc</a></li>
			</ul>
		</div>
{{end}}