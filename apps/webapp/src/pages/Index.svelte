<script>
  import { link, navigate } from "svelte5-router";
  import { getDayPlan } from "../api"
  import { profile } from "../shared"
  import Navbar from "../components/Navbar.svelte";
  import Footer from "../components/Footer.svelte";
  import { onMount } from "svelte";


  if ($profile.name == "") {
    navigate("/profile", {replace: true})
  }

  const now = new Date();
  let dayPlan = $state(null)
  let currTime = $state((now.getHours() * 60) + now.getMinutes())


  onMount(async () => {
    dayPlan = await getDayPlan()

    setInterval(() => {

      currTime = (now.getHours() * 60) + now.getMinutes()
    }, 5000)
  })


  const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
  const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

  const dayName = days[now.getDay()];
  const monthName = months[now.getMonth()];

</script>

<Navbar/>

<main class="container">
  <!-- {JSON.stringify($profile)} -->
  
  <h2>{`${dayName}, ${monthName} ${now.getDate()}, ${now.getFullYear()} ${now.getHours()}:${now.getMinutes()}`}</h2>

  <div class="day_description">
    {dayPlan?.description}
  </div>

  <div class="myday">
      <div class="myday__hour">00:00</div>
      <div class="myday__hour">01:00</div>
      <div class="myday__hour">02:00</div>
      <div class="myday__hour">03:00</div>
      <div class="myday__hour">04:00</div>
      <div class="myday__hour">05:00</div>
      <div class="myday__hour">06:00</div>
      <div class="myday__hour">07:00</div>
      <div class="myday__hour">08:00</div>
      <div class="myday__hour">09:00</div>
      <div class="myday__hour">10:00</div>
      <div class="myday__hour">11:00</div>
      <div class="myday__hour">12:00</div>
      <div class="myday__hour">13:00</div>
      <div class="myday__hour">14:00</div>
      <div class="myday__hour">15:00</div>
      <div class="myday__hour">16:00</div>
      <div class="myday__hour">17:00</div>
      <div class="myday__hour">18:00</div>
      <div class="myday__hour">19:00</div>
      <div class="myday__hour">20:00</div>
      <div class="myday__hour">21:00</div>
      <div class="myday__hour">22:00</div>
      <div class="myday__hour">23:00</div>

      {#if dayPlan?.wakeupTime}
        <div class="myday__wakeup" style="--time: {dayPlan.wakeupTime}; --duration: 10;"><svg><use href="#fa-alarmclock" /></svg></div>
      {/if}
      
      {#if dayPlan?.sleepTime}
        <div class="myday__sleeptime" style="--time: {dayPlan.sleepTime}; --duration: 10;"><svg><use href="#fa-bed" /></svg></div>
      {/if}


      <!-- {#if dayPlan.date == ""} -->
      <div class="myday__now" style="--time: {currTime}; --duration: 3;"></div>

      {#each dayPlan?.events as event }
        <div class="myday__event" style="--time: {event.start}; --duration: {event.duration};">
          <div class="myday__icon">
            {#if event.type == "gym"}
              <svg><use href="#fa-dumbbell" /></svg>
            {:else if event.type == "run"}
              <svg><use href="#fa-person-running" /></svg>
            {:else if event.type == "work"}
              <svg><use href="#fa-person-digging" /></svg>
            {:else if event.type == "meeting"}
              <svg><use href="#fa-handshake" /></svg>
            {:else if event.type == "work"}
              <svg><use href="#fa-person-digging" /></svg>
            {:else if event.type == "travel"}
              <svg><use href="#fa-car" /></svg>
            {:else}
              <svg><use href="#fa-calendar" /></svg>
            {/if}
          </div>
          <div class="myday__eventname">
            {event.title}
          </div>
          <div class="myday__eventdescription">
            {event.description}
          </div>
        </div>
      {/each}

      {#each dayPlan?.suggestions as suggestion }
        <div class="myday__suggestion" style="--time: {suggestion.start}; --duration: {suggestion.duration};">
          <div class="myday__icon">
            {#if suggestion.type == "coffee"}
              <svg><use href="#fa-mug-hot" /></svg>
            {:else if suggestion.type == "sleepTime"}
              <svg><use href="#fa-bed" /></svg>
            {:else if suggestion.type == "pills"}
              <svg><use href="#fa-pills" /></svg>
            {:else if suggestion.type == "food"}
              <svg><use href="#fa-utensils" /></svg>
            {:else if suggestion.type == "glasses"}
              <svg><use href="#fa-glasses" /></svg>
            {:else}
              <svg><use href="#fa-smile" /></svg>
            {/if}
          </div>
          <div class="myday__eventname">
            {suggestion.title}
          </div>
          <div class="myday__eventdescription">
            {suggestion.description}
          </div>
        </div>

      {/each}

  </div>
</main>

<svg xmlns="http://www.w3.org/2000/svg" style="display: none;">

  <symbol id="fa-handshake" viewBox="0 0 576 512">
    <path d="M268.9 85.2L152.3 214.8c-4.6 5.1-4.4 13 .5 17.9 30.5 30.5 80 30.5 110.5 0l31.8-31.8c4.2-4.2 9.5-6.5 14.9-6.9 6.8-.6 13.8 1.7 19 6.9L505.6 376 576 320 576 32 464 96 440.2 80.1C424.4 69.6 405.9 64 386.9 64l-70.4 0c-1.1 0-2.3 0-3.4 .1-16.9 .9-32.8 8.5-44.2 21.1zM116.6 182.7L223.4 64 183.8 64c-25.5 0-49.9 10.1-67.9 28.1L112 96 0 32 0 320 156.4 450.3c23 19.2 52 29.7 81.9 29.7l15.7 0-7-7c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l41 41 9 0c19.1 0 37.8-4.3 54.8-12.3L359 441c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l32 32 17.5-17.5c8.9-8.9 11.5-21.8 7.6-33.1l-137.9-136.8-14.9 14.9c-49.3 49.3-129.1 49.3-178.4 0-23-23-23.9-59.9-2.2-84z"/>
  </symbol>

  <symbol id="fa-person-digging" viewBox="0 0 576 512">
    <path d="M208 40a56 56 0 1 1 112 0 56 56 0 1 1 -112 0zM10.5 181.3c5.9-11.9 20.3-16.7 32.2-10.7l24.6 12.3 12.2-20.4c18.9-31.5 53.2-50.5 89.6-50.5 46.2 0 87.7 30.5 100.5 75.4l32.2 112.7 92.9 46.4 25.8-43c5.8-9.6 16.2-15.5 27.4-15.5s21.7 5.9 27.4 15.5l96 160c5.9 9.9 6.1 22.2 .4 32.2S555.5 512 544 512l-192 0c-11.5 0-22.2-6.2-27.8-16.2s-5.5-22.3 .4-32.2L370 387.8 21.3 213.5c-11.9-5.9-16.7-20.3-10.7-32.2zM94.3 307.4l112 56c10.8 5.4 17.7 16.5 17.7 28.6l0 88c0 17.7-14.3 32-32 32s-32-14.3-32-32l0-68.2-61.3-30.7-36.3 109c-5.6 16.8-23.7 25.8-40.5 20.2S-3.9 486.6 1.7 469.9l48-144c2.9-8.8 9.5-15.9 18.1-19.4s18.3-3.2 26.6 .9z"/>
  </symbol>

  <symbol id="fa-alarmclock" viewBox="0 0 512 512">
    <path d="M504.4 132.5c-4.5 10.5-18.4 9.8-24.9 .4-27.8-40-66.1-72.2-111-92.6-10.4-4.7-13.7-18.3-4.1-24.6 15-9.9 33-15.7 52.3-15.7 52.6 0 95.2 42.6 95.2 95.2 0 13.2-2.7 25.8-7.6 37.3zm-471.9 .4c-6.5 9.4-20.5 10.1-24.9-.4-4.9-11.5-7.6-24.1-7.6-37.3 0-52.6 42.6-95.2 95.2-95.2 19.3 0 37.3 5.8 52.3 15.7 9.6 6.3 6.3 19.9-4.1 24.6-44.8 20.4-83.1 52.6-111 92.6zM390.2 467.4C352.8 495.4 306.3 512 256 512s-96.8-16.6-134.1-44.6L86.6 502.6c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3l35.2-35.2C48.6 384.8 32 338.3 32 288 32 164.3 132.3 64 256 64S480 164.3 480 288c0 50.3-16.6 96.8-44.6 134.2l35.2 35.2c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0l-35.2-35.2zM280 184c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 104c0 6.4 2.5 12.5 7 17l56 56c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-49-49 0-94.1z"/>
  </symbol>

  <symbol id="fa-bed" viewBox="0 0 576 512">
    <path d="M32 32c17.7 0 32 14.3 32 32l0 224 224 0 0-128c0-17.7 14.3-32 32-32l160 0c53 0 96 43 96 96l0 224c0 17.7-14.3 32-32 32s-32-14.3-32-32l0-64-448 0 0 64c0 17.7-14.3 32-32 32S0 465.7 0 448L0 64C0 46.3 14.3 32 32 32zm80 160a64 64 0 1 1 128 0 64 64 0 1 1 -128 0z"/>
  </symbol>

  <symbol id="fa-dumbbell" viewBox="0 0 640 512">
    <path d="M96 112c0-26.5 21.5-48 48-48s48 21.5 48 48l0 112 256 0 0-112c0-26.5 21.5-48 48-48s48 21.5 48 48l0 16 16 0c26.5 0 48 21.5 48 48l0 48c17.7 0 32 14.3 32 32s-14.3 32-32 32l0 48c0 26.5-21.5 48-48 48l-16 0 0 16c0 26.5-21.5 48-48 48s-48-21.5-48-48l0-112-256 0 0 112c0 26.5-21.5 48-48 48s-48-21.5-48-48l0-16-16 0c-26.5 0-48-21.5-48-48l0-48c-17.7 0-32-14.3-32-32s14.3-32 32-32l0-48c0-26.5 21.5-48 48-48l16 0 0-16z"/>
  </symbol>

  <symbol id="fa-person-running" viewBox="0 0 640 512">
    <path d="M256.5-32a56 56 0 1 1 0 112 56 56 0 1 1 0-112zM123.6 176c-3.3 0-6.2 2-7.4 5L94.2 235.9c-6.6 16.4-25.2 24.4-41.6 17.8s-24.4-25.2-17.8-41.6l21.9-54.9C67.7 129.9 94.1 112 123.6 112l97.3 0c28.5 0 54.8 15.1 69.1 39.7l32.8 56.3 61.6 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-61.6 0c-22.8 0-43.8-12.1-55.3-31.8l-10-17.1-20.7 70.4 75.4 22.6c27.7 8.3 41.8 39 30.1 65.5L285.7 509c-7.2 16.2-26.1 23.4-42.2 16.2s-23.4-26.1-16.2-42.2l49.2-110.8-95.9-28.8c-32.7-9.8-52-43.7-43.7-76.8l22.7-90.6-35.9 0zm-8 181c13.3 14.9 30.7 26.3 51.2 32.4l4.7 1.4-6.9 19.3c-5.8 16.3-16 30.8-29.3 41.8L52.9 519.8c-13.6 11.2-33.8 9.3-45-4.3s-9.3-33.8 4.3-45l82.4-67.9c4.5-3.7 7.8-8.5 9.8-13.9L115.6 357z"/>
  </symbol>

  <symbol id="fa-mug-hot" viewBox="0 0 576 512">
    <path d="M152-16c-13.3 0-24 10.7-24 24 0 38.9 23.4 59.4 39.1 73.1l1.1 1c16.3 14.3 23.8 21.8 23.8 37.9 0 13.3 10.7 24 24 24s24-10.7 24-24c0-38.9-23.4-59.4-39.1-73.1l-1.1-1C183.5 31.7 176 24.1 176 8 176-5.3 165.3-16 152-16zM96 192c-17.7 0-32 14.3-32 32l0 192c0 53 43 96 96 96l192 0c41.8 0 77.4-26.7 90.5-64l5.5 0c70.7 0 128-57.3 128-128S518.7 192 448 192L96 192zM448 384l0-128c35.3 0 64 28.7 64 64s-28.7 64-64 64zM288 8c0-13.3-10.7-24-24-24S240-5.3 240 8c0 38.9 23.4 59.4 39.1 73.1l1.1 1c16.3 14.3 23.8 21.8 23.8 37.9 0 13.3 10.7 24 24 24s24-10.7 24-24c0-38.9-23.4-59.4-39.1-73.1l-1.1-1C295.5 31.7 288 24.1 288 8z"/>
  </symbol>

  <symbol id="fa-pills" viewBox="0 0 512 512">
    <path d="M64 112c0-26.5 21.5-48 48-48s48 21.5 48 48l0 112-96 0 0-112zM176 368c0-48.7 18.1-93.2 48-127l0-129C224 50.1 173.9 0 112 0S0 50.1 0 112L0 400c0 61.9 50.1 112 112 112 37.3 0 70.3-18.2 90.7-46.3-17-28.6-26.7-62-26.7-97.7zm64.7 67.4c4.6 8.7 16.3 9.7 23.3 2.7L438.1 264c7-7 6-18.7-2.7-23.3-20.1-10.7-43-16.7-67.4-16.7-79.5 0-144 64.5-144 144 0 24.3 6 47.3 16.7 67.4zM297.9 472c-7 7-6 18.7 2.7 23.3 20.1 10.7 43 16.7 67.4 16.7 79.5 0 144-64.5 144-144 0-24.3-6-47.3-16.7-67.4-4.6-8.7-16.3-9.7-23.3-2.7L297.9 472z"/>
  </symbol>

  <symbol id="fa-car" viewBox="0 0 640 512">
    <path d="M147 106.7l-29.8 85.3 122.9 0 0-96-77.9 0c-6.8 0-12.9 4.3-15.1 10.7zM48.6 193.9L86.5 85.6C97.8 53.5 128.1 32 162.1 32L360 32c25.2 0 48.9 11.9 64 32l96.2 128.3C587.1 196.5 640 252.1 640 320l0 16c0 35.3-28.7 64-64 64l-16.4 0c-4 44.9-41.7 80-87.6 80s-83.6-35.1-87.6-80l-144.7 0c-4 44.9-41.7 80-87.6 80s-83.6-35.1-87.6-80l-.4 0c-35.3 0-64-28.7-64-64l0-80c0-30.1 20.7-55.3 48.6-62.1zM440 192l-67.2-89.6c-3-4-7.8-6.4-12.8-6.4l-72 0 0 96 152 0zM152 432a40 40 0 1 0 0-80 40 40 0 1 0 0 80zm360-40a40 40 0 1 0 -80 0 40 40 0 1 0 80 0z"/>
  </symbol>

  <symbol id="fa-utensils" viewBox="0 0 512 512">
    <path d="M63.9 14.4C63.1 6.2 56.2 0 48 0s-15.1 6.2-16 14.3L17.9 149.7c-1.3 6-1.9 12.1-1.9 18.2 0 45.9 35.1 83.6 80 87.7L96 480c0 17.7 14.3 32 32 32s32-14.3 32-32l0-224.4c44.9-4.1 80-41.8 80-87.7 0-6.1-.6-12.2-1.9-18.2L223.9 14.3C223.1 6.2 216.2 0 208 0s-15.1 6.2-15.9 14.4L178.5 149.9c-.6 5.7-5.4 10.1-11.1 10.1-5.8 0-10.6-4.4-11.2-10.2L143.9 14.6C143.2 6.3 136.3 0 128 0s-15.2 6.3-15.9 14.6L99.8 149.8c-.5 5.8-5.4 10.2-11.2 10.2-5.8 0-10.6-4.4-11.1-10.1L63.9 14.4zM448 0C432 0 320 32 320 176l0 112c0 35.3 28.7 64 64 64l32 0 0 128c0 17.7 14.3 32 32 32s32-14.3 32-32l0-448c0-17.7-14.3-32-32-32z"/>
  </symbol>

  <symbol id="fa-calendar" viewBox="0 0 448 512">
    <path d="M128 0C110.3 0 96 14.3 96 32l0 32-32 0C28.7 64 0 92.7 0 128l0 48 448 0 0-48c0-35.3-28.7-64-64-64l-32 0 0-32c0-17.7-14.3-32-32-32s-32 14.3-32 32l0 32-128 0 0-32c0-17.7-14.3-32-32-32zM0 224L0 416c0 35.3 28.7 64 64 64l320 0c35.3 0 64-28.7 64-64l0-192-448 0z"/>
  </symbol>

  <symbol id="fa-smile" viewBox="0 0 512 512">
    <path d="M464 256a208 208 0 1 0 -416 0 208 208 0 1 0 416 0zM0 256a256 256 0 1 1 512 0 256 256 0 1 1 -512 0zm177.3 63.4C192.3 335 218.4 352 256 352s63.7-17 78.7-32.6c9.2-9.6 24.4-9.9 33.9-.7s9.9 24.4 .7 33.9c-22.1 23-60 47.4-113.3 47.4s-91.2-24.4-113.3-47.4c-9.2-9.6-8.9-24.8 .7-33.9s24.8-8.9 33.9 .7zM144 208a32 32 0 1 1 64 0 32 32 0 1 1 -64 0zm192-32a32 32 0 1 1 0 64 32 32 0 1 1 0-64z"/>
  </symbol>

  <symbol id="fa-glasses" viewBox="0 0 576 512">
    <path d="M143.3 96c-14 0-26.5 9.2-30.6 22.6L70.4 256 224 256c17.7 0 32 14.3 32 32l64 0c0-17.7 14.3-32 32-32l153.6 0-42.3-137.4C459.2 105.2 446.8 96 432.7 96L400 96c-17.7 0-32-14.3-32-32s14.3-32 32-32l32.7 0c42.1 0 79.4 27.5 91.8 67.8l45.4 147.5c4.1 13.2 6.1 26.9 6.1 40.7l0 96c0 53-43 96-96 96l-64 0c-53 0-96-43-96-96l0-32-64 0 0 32c0 53-43 96-96 96l-64 0c-53 0-96-43-96-96l0-96c0-13.8 2.1-27.5 6.1-40.7L51.5 99.8C63.9 59.5 101.1 32 143.3 32L176 32c17.7 0 32 14.3 32 32s-14.3 32-32 32l-32.7 0zM64 320l0 64c0 17.7 14.3 32 32 32l64 0c17.7 0 32-14.3 32-32l0-64-128 0zm416 96c17.7 0 32-14.3 32-32l0-64-128 0 0 64c0 17.7 14.3 32 32 32l64 0z"/>
  </symbol>
</svg>

<Footer/>
  
<style lang="scss">
  main {
    margin-top: 60px;
  }

  .day_description {
    background-color: #ddd;
    padding: 20px;
    border-radius: 20px;
    margin: 20px 0;
  }
  
  .myday {
    position: relative;

    

    &__hour {
      height: 60px;
      line-height: 60px;
      font-size: 1rem;
      border-bottom: 1px solid #ccc;
    }

    &__eventname {
      font-size: 1rem;
      color: #222;
      font-weight: 600;
    }

    &__eventdescription {
      font-size: 0.8rem;
      color: #ccc;
    }

    &__wakeup,
    &__sleeptime {
      margin-left: 50px;
      svg {
        height: 24px;
        max-width: 24px;
        fill: #153ac0;
      }
    }

    &__now {
      z-index: 11;
      opacity: 0.7;
      background-color: #7615c0;
      margin-left: 45px;
      width: calc(100% - 45px);
    }

    &__suggestion,
    &__event {
      margin-left: 90px;
      background-color: #ccc;
      border-radius: 10px;
      padding: 10px;
      display: flex;
      flex-direction: row;
      width: calc(100% - 100px);
      opacity: 0.95;

      :hover {
        opacity: 1;
      }
    }


    &__suggestion {
      margin-left: 55px;
      width: calc(100% - 55px);
      background-color: #1b09bd;
      padding: 4px;
      min-height: 30px;
    }

    &__wakeup,
    &__sleeptime,
    &__now,
    &__suggestion,
    &__event {
      position: absolute;
      top: calc(var(--time) * 1px);
      height: calc(var(--duration) * 1px);
    }
    

    &__icon {
      width: 40px;
      svg {
        height: 32px;
        max-width: 32px;
      }
    }

  }

</style>
