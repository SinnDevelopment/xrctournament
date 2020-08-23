package main

// TODO: Hardcode files in here.
func getData(name string) string {
	switch name {
	case "index.html":
		return `<!doctype html>
		<html lang="en">
		
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<meta name="description" content="">
			<meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
			<meta name="generator" content="Jekyll v4.1.1">
			<title>{{.TConf.CompetitionName}}</title>
		
		
			<!-- Bootstrap core CSS -->
			<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"
				integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj"
				crossorigin="anonymous"></script>
			<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js"
				integrity="sha384-9/reFTGAW83EW2RDu2S0VKaIzap3H66lZH81PoYlFhbGU+6BZp6G7niu735Sk7lN"
				crossorigin="anonymous"></script>
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css"
				integrity="sha384-JcKb8q3iqJ61gNV9KGb8thSsNjpSL0n8PARn9HuZOnIxN0hoP+VmmDGMN5t9UJ0Z" crossorigin="anonymous">
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"
				integrity="sha384-B4gt1jrGC7Jh4AgTPSdUtOBvfO8shuf57BaghqFfPlYxofvL8/KUEfYiJOMMV+rV"
				crossorigin="anonymous"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/feather-icons/4.9.0/feather.min.js"></script>
		
		
			<style>
				.bd-placeholder-img {
					font-size: 1.125rem;
					text-anchor: middle;
					-webkit-user-select: none;
					-moz-user-select: none;
					-ms-user-select: none;
					user-select: none;
				}
		
				@media (min-width: 768px) {
					.bd-placeholder-img-lg {
						font-size: 3.5rem;
					}
				}
		
				body {
					font-size: .875rem;
				}
		
				.feather {
					width: 16px;
					height: 16px;
					vertical-align: text-bottom;
				}
		
				/*
		 * Sidebar
		 */
		
				.sidebar {
					position: fixed;
					top: 0;
					bottom: 0;
					left: 0;
					z-index: 100;
					/* Behind the navbar */
					padding: 48px 0 0;
					/* Height of navbar */
					box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);
				}
		
				@media (max-width: 767.98px) {
					.sidebar {
						top: 5rem;
					}
				}
		
				.sidebar-sticky {
					position: relative;
					top: 0;
					height: calc(100vh - 48px);
					padding-top: .5rem;
					overflow-x: hidden;
					overflow-y: auto;
					/* Scrollable contents if viewport is shorter than content. */
				}
		
				@supports ((position: -webkit-sticky) or (position: sticky)) {
					.sidebar-sticky {
						position: -webkit-sticky;
						position: sticky;
					}
				}
		
				.sidebar .nav-link {
					font-weight: 500;
					color: #333;
				}
		
				.sidebar .nav-link .feather {
					margin-right: 4px;
					color: #999;
				}
		
				.sidebar .nav-link.active {
					color: #007bff;
				}
		
				.sidebar .nav-link:hover .feather,
				.sidebar .nav-link.active .feather {
					color: inherit;
				}
		
				.sidebar-heading {
					font-size: .75rem;
					text-transform: uppercase;
				}
		
				/*
		 * Navbar
		 */
		
				.navbar-brand {
					padding-top: .75rem;
					padding-bottom: .75rem;
					font-size: 1rem;
					background-color: rgba(0, 0, 0, .25);
					box-shadow: inset -1px 0 0 rgba(0, 0, 0, .25);
				}
		
				.navbar .navbar-toggler {
					top: .25rem;
					right: 1rem;
				}
		
				.navbar .form-control {
					padding: .75rem 1rem;
					border-width: 0;
					border-radius: 0;
				}
		
				.form-control-dark {
					color: #fff;
					background-color: rgba(255, 255, 255, .1);
					border-color: rgba(255, 255, 255, .1);
				}
		
				.form-control-dark:focus {
					border-color: transparent;
					box-shadow: 0 0 0 3px rgba(255, 255, 255, .25);
				}
			</style>
		</head>
		
		<body>
			<nav class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
				<a class="navbar-brand col-md-3 col-lg-2 mr-0 px-3" href="#">{{.TConf.CompetitionName}}</a>
				<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-toggle="collapse"
					data-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
				<ul class="navbar-nav px-3">
					<li class="nav-item text-nowrap">
						<a class="nav-link" href="//github.com/SinnDevelopment/xrctournament">GitHub</a>
					</li>
				</ul>
			</nav>
		
			<div class="container-fluid">
				<div class="row">
					<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
						<div class="sidebar-sticky pt-3">
							<ul class="nav flex-column">
								<li class="nav-item">
									<a class="nav-link active" href="/">
										Home
									</a>
								</li>
								<li class="nav-item">
									<a class="nav-link" href="/rankings">
										Rankings
									</a>
								</li>
								<li class="nav-item">
									<a class="nav-link" href="/matches">
										Matches
									</a>
								</li>
							</ul>
						</div>
					</nav>
		
					<main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">
						{{if eq .Page "home" }}
						<div
							class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
							<h1 class="h2">{{.TConf.CompetitionName}}</h1>
						</div>
						<!-- Add a placeholder for the Twitch embed -->
						<div id="twitch-embed"></div>
						<script src="https://embed.twitch.tv/embed/v1.js"></script>
		
						<!-- Create a Twitch.Embed object that will render within the "twitch-embed" root element. -->
						<script type="text/javascript">
							new Twitch.Embed("twitch-embed", {
								width: 854,
								height: 480,
								channel: "{{.TConf.TwitchChannel}}"
							});
						</script>
						{{end}}
						{{if eq .Page "rankings" }}
						<h2>Rankings</h2>
						<div class="table-responsive">
							<table class="table table-striped table-sm">
								<thead>
									<tr>
										<th>#</th>
										<th>Name</th>
										<th>OPR</th>
										<th>Points</th>
										<th>W/L/T</th>
									</tr>
								</thead>
								<tbody>
									{{range .Players}}
									<tr>
										<td>#</td>
										<td>{{.Name}}</td>
										<td>{{.OPR}}/</td>
										<td>{{.Points}}</td>
										<td>{{.Wins}}/{{.Losses}}/{{.Ties}}</td>
									</tr>
									{{end}}
								</tbody>
							</table>
						</div>
						{{end}}
						{{if eq .Page "matches" }}
						<h2>Matches</h2>
						<div class="table-responsive">
							<table class="table table-striped table-sm">
								<thead>
									<tr>
										<th>#</th>
										<th>R1</th>
										<th>R2</th>
										<th>R3</th>
										<th>B1</th>
										<th>B2</th>
										<th>B3</th>
										<th>Red</th>
										<th>Blue</th>
									</tr>
								</thead>
								<tbody>
									{{range .Matches}}
									<tr>
										<td>#</td>
										<td>{{.Red1}}</td>
										<td>{{.Red2}}/</td>
										<td>{{.Red3}}</td>
										<td>{{.Blue1}}</td>
										<td>{{.Blue2}}/</td>
										<td>{{.Blue3}}</td>
										<td>{{.RedScore}}</td>
										<td>{{.BlueScore}}</td>
									</tr>
									{{end}}
								</tbody>
							</table>
						</div>
						{{end}}
					</main>
				</div>
			</div>
		
		
		</body>
		
		</html>`
	default:
		return ""
	}
}
