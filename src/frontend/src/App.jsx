
import {useState, useEffect } from 'react'
import './index.css';
import Header from './components/Header'
import Search from './components/Search'

const App = () => {
  const[loadStatus, setLoadStatus] = useState(false)

  function retrieveSummoner(userName) {
    fetch(`http://localhost:8080/player?name=${userName}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    })
    .then((response) => response.json())
    .then((responseJSON) => {
      if (responseJSON) {
        console.log(responseJSON)
      }
    })
    .catch((err)=> {
      console.log(err)
    });
  }

  return (
    <div className="container">
      <Header/>
      <Search getPlayerInformation = {retrieveSummoner}/>
    </div>
  );
}

export default App;
