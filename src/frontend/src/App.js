
import {useState, useEffect } from 'react'
import './index.css';
import Header from './components/Header'
import Search from './components/Search'

const App = () => {
  const[items, setItems] = useState([])
  const[loadStatus, setLoadStatus] = useState(false)

  const getValue = (event) => {
    console.log('Event: ', event)
  }

  const componentDidMount = (playerName) => e => {
    fetch(`localhost:8080/player/${playerName}`)
    .then((response) => {
      if (window.status === 200) {
        return response.json();
      }
      throw new Error('Error: Invalid Summoner Name');
    })
    .then((responseJSON) => {
      console.log(responseJSON)
    })
    .catch((err)=> {
      throw new Error('Error: Invalid Summoner Name')
    });
  }

  return (
    <div className="container">
      <Header/>
      <Search getValueFunctionFromMain = {getValue}/>
    </div>
  );
}

export default App;
