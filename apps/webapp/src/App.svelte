<script>
	import { Router, Route } from "svelte5-router";
  import Index from "./pages/Index.svelte"
  import Profile from "./pages/Profile.svelte"
</script>

<Router>
	<Route path="/"><Index /></Route>
	<Route path="/profile"><Profile /></Route>
</Router>

<style>

</style>
