<script>
	import { Router, Route } from "svelte5-router";
  import Index from "./pages/Index.svelte"
  import Profile from "./pages/Profile.svelte"

	let { url } = $props();
	console.log("url", url)
</script>

<Router {url}>
	<Route path="/"><Index /></Route>
	<Route path="/profile"><Profile /></Route>
</Router>

<style>

</style>
