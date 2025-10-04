<script>
    import Navbar from "../components/Navbar.svelte";
    import { profile, enableNotifications } from "../shared";
	import { link } from "svelte5-router";
    import { subscribeNotifications } from '../api'

    let profileData = $state({
        name: "",
        calendarUrl: "",
        ...$profile
    })

    const saveProfile = (e) => {
        console.log('saveProfile', profileData)

        profile.set(profileData)
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
                    <label for="name" class="form-label">Imię</label>
                    <input type="text" id="name" class="form-control" placeholder="Imię" bind:value={profileData.name}>
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
                    <label for="calendarUrl" class="form-label">Link do kalendarza (iCal)</label>
                    <input type="text" id="calendarUrl" class="form-control" placeholder="Link do kalendarza (iCal)" bind:value={profileData.calendarUrl}>
                </div>

                <div class="form-group">
                    <label for="notifications" class="form-label">Powiadomienia</label>
                    <button type="button" onclick={enableNotifications} class="button button-primary">Włącz powiadomienia</button>
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
        }
    </style>