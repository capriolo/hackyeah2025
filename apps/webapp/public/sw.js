// sw.js
self.addEventListener('push', event => {
  console.log('[Service Worker] Push Received');
  console.log('[Service Worker] Push had this data:', event.data);

  try {
    const data = event.data.json();
    console.log('[Service Worker] Parsed data:', data);
    
    event.waitUntil(
      self.registration.showNotification(data.title, {
        body: data.body,
        icon: "/icon.png",
        badge: "/badge.png",
        vibrate: [200, 100, 200]
      }).then(() => {
        console.log('[Service Worker] Notification shown successfully');
      }).catch(err => {
        console.error('[Service Worker] Notification error:', err);
      })
    );
  } catch (error) {
    console.error('[Service Worker] Error parsing push data:', error);
    // Fallback dla prostego tekstu
    const title = 'New Notification';
    const options = {
      body: event.data ? event.data.text() : 'No content',
      icon: "/icon.png"
    };
    event.waitUntil(self.registration.showNotification(title, options));
  }
});

self.addEventListener('notificationclick', event => {
  console.log('[Service Worker] Notification click Received');
  event.notification.close();
  event.waitUntil(
    clients.openWindow('https://localhost')
  );
});

// Dodaj te event listeners dla debugowania
self.addEventListener('install', event => {
  console.log('[Service Worker] Installed');
  self.skipWaiting();
});

self.addEventListener('activate', event => {
  console.log('[Service Worker] Activated');
  event.waitUntil(clients.claim());
});