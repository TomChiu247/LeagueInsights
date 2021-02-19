import{ useState } from 'react'

const Search = ({getPlayerInformation}) => {
    const[userName, setUserName] = useState('') 

    const handleSubmit = (e) => {
        e.preventDefault();
        getPlayerInformation(userName);
    }

    return (
        <div>
        <form className='add-form' onSubmit={handleSubmit}>
        <div className='form-control'>
            <label htmlFor=""></label>
            <input 
                type='text'
                placeholders='Enter a summoner name'
                onChange={(e)=>setUserName(e.target.value)}
            />
        </div>  
        <input type='submit' value='Search' className='btn btn-block' />
        </form>
    </div>
        
    )
}

export default Search
