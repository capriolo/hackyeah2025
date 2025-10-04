import { writable } from 'svelte/store';

const createLocalStorageStore = (key, initialValue) => {
    if (!import.meta.env.SSR) {
        const storedValue = localStorage.getItem(key);
        initialValue = storedValue ? JSON.parse(storedValue) : initialValue;
    }
    
    const store = writable(initialValue);

    if (!import.meta.env.SSR) {
        store.subscribe(val => {
            if (val !== undefined) {
                localStorage.setItem(key, JSON.stringify(val));
            }
        });
    }

    return store;
}

const profile = createLocalStorageStore('profile', {
    name: "",
    age: 10,
});

const logout = () => {
    profile.set(null)
}

export { profile, logout }