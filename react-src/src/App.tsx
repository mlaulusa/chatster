import React, { useContext } from 'react'
import map from 'lodash/map'
import './App.css'
import { AppContext } from './state/Provider'
import ChatRoom from './ChatRoom'

interface props {

}

const App: React.FunctionComponent<props> = () => {
  const { state: { rooms } } = useContext(AppContext)

  return (
    <>
      {map(rooms, room => <ChatRoom name={room} />)}
    </>
  )
}

export default App
