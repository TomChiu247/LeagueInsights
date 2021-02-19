
import {useState, useEffect } from 'react'
import './index.css';
import Header from './components/Header'
import Search from './components/Search'

const App = () => {
  //const[items, setItems] = useState([])
  const[loadStatus, setLoadStatus] = useState(false)

  //const getValue = (playerName) => {
  //  setUserName(`${playerName}`)
  //  console.log(userName)
  //}
  //'https://agile-chamber-72618.herokuapp.com/http://localhost:8080/player?name=${userName}`
  function retrieveSummoner(userName) {
    fetch(`http://localhost:8080/player?name=${userName}`, {
      //mode: 'no-cors',
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    })
    .then((response) => {
      if (response.status === 200) {
        return response.json();
      }
      else{
        throw new Error('Error: Invalid Summoner Name');
      }
    })
    .then((responseJSON) => {
      console.log(responseJSON)
    })
    .catch((err)=> {
      console.log('Error: Invalid Summoner Name')
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
