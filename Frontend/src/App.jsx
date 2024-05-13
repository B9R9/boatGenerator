import { useEffect, useState } from 'react'

import BoatsServices from './services/boats'
import SettingsServices from './services/settings'

import DisplayBoats from './components/DisplayBoats'
import DisplaySettings from './components/DisplaySettings';

function App() {
  
  const [boats, setBoats] = useState([]); // Initialiser boats à un tableau vide
  const [settings, setSettings] = useState(null)
  const [refreshRate, setRefreshRate] = useState(0)
  
  useEffect(() => {
      console.log("Effect")
      const fetchData = async () => {
        const res = await SettingsServices.get()
          setSettings(res)
          setRefreshRate(res.RefreshRate)
    }
    fetchData()
    .catch(console.error)
  }, [])

  useEffect(() => {
    const fetchBoats = async () => {
      const res = await BoatsServices.getAll()
      setBoats(res)
    }
    fetchBoats()

    console.log("Refreshrate --->", refreshRate)
    const intervalId = setInterval(() => {
      fetchBoats();
    }, refreshRate) 

    // Nettoyer l'intervalle lorsque le composant est démonté
    return () => clearInterval(intervalId);
  }, [refreshRate]); // Effectuer le fetch de bateaux chaque fois que la fréquence de rafraîchissement change

  const handleRefreshRateChange = (value) => {
    setRefreshRate(value);
  }

  return (
    <>
    <DisplaySettings settings={settings} onChange={handleRefreshRateChange}/>
    <DisplayBoats boats={boats} />
    </>
  ) 
}

export default App
