{{define "navbar"}}
		<a class="navbar-brand" href="/">GoChair</a>
		<div>
			<ul class="nav navbar-nav">
				<li {{if .IsIndex}}class="active"{{end}}><a href="/">Chair</a></li>
				<li {{if .IsProject}}class="active"{{end}}><a href="/project">Project</a></li>
				<li {{if .IsTask}}class="active"{{end}}><a href="/task">Task</a></li>
			</ul>
		</div>
{{end}}