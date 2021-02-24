
const Button = ({getValueFunction}) => {
    return (
            <button 
                onClick = {getValueFunction} 
                className='btn btn-block'>
                Submit
            </button>
    )
}

export default Button
