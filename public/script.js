class ApiService {
    async shortenURL(url) {
        const rsp = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({url: url})
        })

        if (!rsp.ok) {
            return {
                error: "Failed to shorten the URL"
            }
        }

        return rsp.json()
    }

    async deleteURL(code, password) {
        const rsp = await fetch(`/${code}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({password})
        })

        if (!rsp.ok) {
            return {
                error: "Failed to delete the URL"
            }
        }

        return rsp.json()
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const api = new ApiService()

    let rspData

    document.querySelector('#shorten-btn').addEventListener('click', async () => {
        const url = document.querySelector('#url-input').value
        const rsp = await api.shortenURL(url)

        if (rsp.error) {
            alert(rsp.error)
            return
        }

        rspData = rsp
        const a = document.querySelector('#result-div-a')
        a.setAttribute('href', rsp.url)
        a.innerHTML = rsp.url

        document.querySelector('#result-div').style.display = "block"
    })

    document.querySelector('#delete-btn').addEventListener('click', async () => {
        if (!rspData) {
            return
        }

        const rsp = await api.deleteURL(rspData.hash, rspData.password)
        if (rsp.error) {
            alert(rsp.error)
            return
        }

        document.querySelector('#result-div').style.display = "none"
    })
})