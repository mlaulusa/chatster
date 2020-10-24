import { useState, useEffect, useContext } from 'react'
import { AppContext } from '../state/Provider'

export default function useFetchRooms () {
  const [pending, setPending] = useState(false)
  const { dispatch } = useContext(AppContext)

  setPending(true)

  fetch("/room", { method: "GET" })
    .then(response => {
      if (response.ok) {
        return response.json()
      } else {
        throw Error("Fetch to '/room' failed")
      }
    })
    .then(data => {
      console.log('data', data)
    })
    .catch(err => {
      console.error(err)
    })
    .finally(() => {
      setPending(false)
    })

  return pending
}