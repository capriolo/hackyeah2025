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
    })

    const garminSync = async (e) => {

        const button = e.currentTarget;
        button.classList.add('loading');

        await garminSyncApiCall()

        button.classList.remove('loading');
    }

    const googleCalSync = async (e) => {
        const button = e.currentTarget;
        button.classList.add('loading');

        await icalSyncApiCall()

        button.classList.remove('loading');
    }
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
                    position: relative;

                    img {
                        max-width: 64px;
                        max-height: 50px;
                    }
                }
                :global(button.loading) {

                    &:after {
                        position: absolute;
                        display: block;
                        top: 0;
                        left: 0;
                        content: '...';
                        width: 100px;
                        height: 64px;
                        background-color: rgba(#fff, 0.6);
                    }
                }
            }
        }
    </style>