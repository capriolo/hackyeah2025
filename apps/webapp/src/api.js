
export const subscribeNotifications = async (sub) => {
    const res = await _post('/api/v1/push', sub)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return Promise.resolve(res)
}

export const saveProfile = async (data) => {
    const res = await _post('/api/v1/profile', data)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return Promise.resolve(res)
}

export const garminSync = async (data) => {
    const res = await _post('/api/v1/sync/garmin', data)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return Promise.resolve(res)
}

export const icalSync = async (data) => {
    const res = await _post('/api/v1/sync/ical', data)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return Promise.resolve(res)
}

export const getProfile = async () => {
    const res = await _get('/api/v1/profile')
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return res.json()
}


export const getDayPlan = async (date) => {
    const res = await _get(`/api/v1/day?date=${date}`)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    
    return res.json()
}

const _get = async (endpoint) => {
    if (import.meta.env.SSR) {
        return Promise.resolve(Response.json({}))
    }

    return await fetch(`${import.meta.env.VITE_API_URL}${endpoint}`, {
        credentials: "include",
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }
    })
}

const _post = async (endpoint, data) => {  
    let headers = {
        'Accept': 'application/json',
        'X-CSRF-Token': globalThis.csrfToken
    }

    if (!(data instanceof FormData)) {
        headers['Content-Type'] = 'application/json'
        data = JSON.stringify(data)
    }

    return await fetch(`${import.meta.env.VITE_API_URL}${endpoint}`, {
        credentials: "include",
        method: 'POST',
        headers: headers,
        body: data
    })
}
