{{define "navbar"}}
		<a class="navbar-brand" href="/">GoChair</a>
		<div>
			<ul class="nav navbar-nav">
				<li {{if .IsHome}}class="active"{{end}}><a href="/">Chair</a></li>
				<li {{if .IsProjcet}}class="active"{{end}}><a href="/Project">Project</a></li>
				<li {{if .IsCommand}}class="active"{{end}}><a href="/Command">Command</a></li>
				<li {{if .IsLoginfo}}class="active"{{end}}><a href="/Loginfo">Loginfo</a></li>
				<li {{if .IsTask}}class="active"{{end}}><a href="/Task">Task</a></li>
			</ul>
		</div>
{{end}}