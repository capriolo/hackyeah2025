<script>
    import Navbar from "../components/Navbar.svelte";
    import { toggleNotifications, hasNotificationsEnabled as hasNotificationsEnabledCall } from "../shared";
	import { link } from "svelte5-router";
    import { getProfile, saveProfile as saveProfileApiCall, garminSync as garminSyncApiCall, icalSync as icalSyncApiCall } from '../api'
    import { onMount } from "svelte";

    let profileData = $state({
        name: "",
        calGoogle: false,
        garmin: false,
    })
    
    const saveProfile = async (e) => {
        e.preventDefault()

        await saveProfileApiCall(profileData)
    }

    let hasNotificationsEnabled = $state(true)

    onMount(async () => {
        hasNotificationsEnabled = await hasNotificationsEnabledCall()
        profileData = await getProfile()

        console.log("profileData", profileData)
    })

    const garminSync = async () => {
        await garminSyncApiCall()
    }

    const googleCalSync = async () => {
        await icalSyncApiCall()
    }

    // let bmr = $derived(10*profileData.weight + 6.25)
    // BMR = 10 × masa(kg) + 6.25 × wzrost(cm) – 5 × wiek(lata) + 5
    // Kobiety:
    // BMR = 10 × masa(kg) + 6.25 × wzrost(cm) – 5 × wiek(lata) – 161

</script>

<Navbar/>

<div class="container">

    <div class="form">
            <form onsubmit={saveProfile}>
                <div>
                    <h3>Twoje dane</h3>
                </div>

                <div class="form-group">
                    <label for="name" class="form-label">Nazwa użytkownika</label>
                    <input type="text" id="name" class="form-control" placeholder="Nazwa użytkownika" bind:value={profileData.name}>
                </div>

                <!-- <div class="form-group">
                    <label for="birthYear" class="form-label">Rok urodzenia</label>
                    <input type="number" id="birthYear" class="form-control" placeholder="Wiek" bind:value={profileData.birthYear}>
                </div>

                <div class="form-group">
                    <label for="name" class="form-label">Płeć</label>
                    <input type="radio" id="name" class="form-control" value="female" bind:group={profileData.gender}> Kobieta
                    <input type="radio" id="name" class="form-control" value="male" bind:group={profileData.gender}> Mężczyzna
                </div> -->

                <!-- <div class="form-group">
                    <label for="weight" class="form-label">Masa ciała [kg]</label>
                    <input type="text" id="weight" class="form-control" placeholder="Masa ciała" bind:value={profileData.weight}>
                </div>

                <div class="form-group">
                    <label for="pal" class="form-label">Poziom aktywności</label> -->

                    <!-- Poziom aktywności	Opis	Współczynnik
                    1.2	Siedzący tryb życia (mało ruchu, praca biurowa)	1.2
                    1.375	Lekka aktywność (ćwiczenia 1–3×/tydz.)	1.375
                    1.55	Średnia aktywność (ćwiczenia 3–5×/tydz.)	1.55
                    1.725	Wysoka aktywność (ćwiczenia 6–7×/tydz.)	1.725
                    1.9	Bardzo wysoka aktywność (ciężka praca fizyczna, sport zawodowy)	1.9 -->

                    <!-- <input type="range" min="1.2" max="1.9" step="0.05" id="pal" class="form-control" bind:value={profileData.pal}>
                </div>
            
                <div class="form-group">
                    <label for="tdee" class="form-label">Zapotrzebowanie kaloryczne</label>
                    <input type="date" id="tdee" class="form-control" placeholder="Wiek" bind:value={profileData.tdee}>
                </div> -->

                <div class="form-group">
                    <label for="calendarUrl" class="form-label">Zsynchronizuj kalendarz</label>

                    <div class="platforms">
                        <button type="button" onclick={googleCalSync}>
                            <img src="/img/cal-google.png" alt="google calendar"/>
                        </button>

                        <button type="button" onclick={()=>{}}>
                            <img src="/img/cal-apple.png" alt="apple calendar"/>
                        </button>

                        <button type="button" onclick={()=>{}}>
                            <img src="/img/cal-ms.png" alt="microsoft calendar"/>
                        </button>
                    </div>

                </div>

                <div class="form-group smartwatch">
                    <label for="smartwatchSync" class="form-label">Zsynchronizuj smartwatch</label>

                    <div class="platforms">
                        <button type="button" onclick={garminSync}>
                            <img src="/img/garminconnect.png" alt="garmin"/>
                        </button>

                        <button type="button" onclick={()=>{}}>
                            <img src="/img/fitbit.png" alt="fitbit"/>
                        </button>

                        <button type="button" onclick={()=>{}}>
                            <img src="/img/polar.png" alt="polar"/>
                        </button>

                        <button type="button" onclick={()=>{}}>
                            <img src="/img/suunto.png" alt="suunto"/>
                        </button>

                    </div>
                    <!-- <input type="text" id="calendarUrl" class="form-control" placeholder="Link do kalendarza (iCal)" bind:value={profileData.calendarUrl}> -->
                </div>

                <div class="form-group">
                    <label for="notifications" class="form-label">Powiadomienia</label>
                    
                    <button type="button" onclick={() => {toggleNotifications(); hasNotificationsEnabled = !hasNotificationsEnabled}} class="button button-primary">
                        {#if hasNotificationsEnabled}
                            Wyłącz
                        {:else}
                            Włącz powiadomienia
                        {/if}
                    </button>
                </div>

                <div class="form-actions">
                    <a use:link href="/" class="button button-secondary">Anuluj</a>
                    <button type="submit" class="button button-primary">Zapisz</button>
                </div>
            </form>
        </div>
    </div>


    <style lang="scss">

        .form {
            margin-top: 60px;

            h3 {
                margin-bottom: 16px;
            }

            .platforms {
                display: flex;
                gap: 4px;

                button {
                    width: 100px;
                    height: 64px;
                    cursor: pointer;
                    background: rgba(0,0,0,0.1);
                    border: none;

                    img {
                        max-width: 64px;
                        max-height: 50px;
                    }
                }
            }
        }
    </style>