import React, { createContext, useMemo, useReducer } from 'react'
import noop from 'lodash/noop'
import reducer, { AppState, initialState } from './reducer'

export const AppContext = createContext<{ state: AppState, dispatch: Function }>({ state: initialState, dispatch: noop })

const Provider: React.FunctionComponent = ({ children }) => {

  const [state, dispatch] = useReducer(reducer, initialState)

  const providerValue = useMemo(() => ({ state, dispatch }), [state])

  return (
    <AppContext.Provider value={providerValue}>
      {children}
    </AppContext.Provider>
  )
}

export default Provider
