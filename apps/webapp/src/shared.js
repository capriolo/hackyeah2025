import { writable } from 'svelte/store';
import { subscribeNotifications } from './api'


const toggleNotifications = async () => {
    try {
        // Sprawdź czy Service Worker jest wspierany
        if (!('serviceWorker' in navigator)) {
            throw new Error('Service Worker nie jest wspierany');
        }

        // Sprawdź czy Push Manager jest wspierany
        if (!('PushManager' in window)) {
            throw new Error('Push notifications nie są wspierane');
        }

        // Zarejestruj Service Worker
        const registration = await navigator.serviceWorker.register('/sw.js');
        console.log('Service Worker zarejestrowany');

        // Sprawdź czy już subskrybujemy
        let subscription = await registration.pushManager.getSubscription();
        
        if (subscription) {
            console.log('Już subskrybujemy powiadomienia');

            unsub()
            return false
            // return subscription
        }

        // Poproś o pozwolenie
        const permission = await Notification.requestPermission();
        
        if (permission !== 'granted') {
            throw new Error('Użytkownik nie zezwolił na powiadomienia');
        }

        const key = import.meta.env.VITE_VAPID_KEY
        // Subskrybuj do push notifications
        subscription = await registration.pushManager.subscribe({
            userVisibleOnly: true,
            applicationServerKey: key
        });

        console.log('Subskrypcja udana:', subscription);
        await subscribeNotifications(subscription)

        return subscription
    } catch (error) {
        console.error('Błąd subskrypcji:', error);
        alert('Nie udało się włączyć powiadomień: ' + error.message);
    }
}



const hasNotificationsEnabled = async () => {
    try {
        // Sprawdź czy Service Worker jest wspierany
        if (!('serviceWorker' in navigator)) {
            throw new Error('Service Worker nie jest wspierany');
        }

        // Sprawdź czy Push Manager jest wspierany
        if (!('PushManager' in window)) {
            throw new Error('Push notifications nie są wspierane');
        }

        // Zarejestruj Service Worker
        const registration = await navigator.serviceWorker.register('/sw.js');
        console.log('Service Worker zarejestrowany');

        // Sprawdź czy już subskrybujemy
        let subscription = await registration.pushManager.getSubscription();
        
        if (subscription) {
            return true
        }

        return false
    } catch (error) {
        console.error('Błąd subskrypcji:', error);
        alert('Nie udało się włączyć powiadomień: ' + error.message);
        return false
    }
}


const unsub = () => {
    navigator.serviceWorker.ready
        .then(registration => {
            registration.pushManager.getSubscription()
            .then(pushSubscription => {
                if(pushSubscription){
                    //check if user was subscribed with a different key
                    let json = pushSubscription.toJSON();
                    let public_key = json.keys.p256dh;
                    
                    console.log(public_key);
                    
                
                        pushSubscription.unsubscribe().then(successful => {
                            console.log("unsubscribed")
                        }).catch(e => {
                            console.log("unsub failed")
                            // Unsubscription failed
                        })
                    }
            });
        })
}

export { toggleNotifications, hasNotificationsEnabled }