
const Button = ({getValueFunction}) => {
    return (
        <div>
            <form action="">
            <div>
                <label htmlFor=""></label>
                <input type="text" onChange = {getValueFunction}/>
            </div>  
            <button>Submit</button>
        </form>
        </div>
    )
}

export default Button
