import React, { useContext } from 'react'
import map from 'lodash/map'
import './App.css'
import { AppContext } from './state/Provider'
import ChatRoom from './ChatRoom'
import useFetchRooms from './hooks/useFetchRooms'

interface props {

}

const App: React.FunctionComponent<props> = () => {
  const { state: { rooms } } = useContext(AppContext)

  const pending = useFetchRooms()

  return pending
    ? (<p>Loading</p>)
    : (
      <>
        {map(rooms, ({ id, name }) => <ChatRoom key={id} name={name} />)}
      </>
    )
}

export default App
