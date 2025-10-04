<script>
  import { link, navigate } from "svelte5-router";
  import { profile } from "../shared"
    import Navbar from "../components/Navbar.svelte";


  if ($profile.name == "") {
    navigate("/profile")
  }

</script>

<Navbar/>

<div class="container">
  {JSON.stringify($profile)}
  
</div>
