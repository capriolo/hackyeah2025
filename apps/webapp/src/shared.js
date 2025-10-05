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

function addDaysToYYYYMMDD(yyyymmdd, daysToAdd) {
  // Parsowanie daty z formatu 20240115
  const year = parseInt(yyyymmdd.slice(0, 4));
  const month = parseInt(yyyymmdd.slice(4, 6)) - 1; // miesiące 0-11
  const day = parseInt(yyyymmdd.slice(6, 8));
  
  const date = new Date(year, month, day);
  
  // Dodanie dni
  date.setDate(date.getDate() + daysToAdd);
  
  // Formatowanie z powrotem do YYYYMMDD
  const newYear = date.getFullYear();
  const newMonth = String(date.getMonth() + 1).padStart(2, '0');
  const newDay = String(date.getDate()).padStart(2, '0');
  
  return `${newYear}${newMonth}${newDay}`;
}

export { toggleNotifications, hasNotificationsEnabled, addDaysToYYYYMMDD }