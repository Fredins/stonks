import React, { useState } from 'react'


const Test = () => {

    const fetchApi = async () => {
        const url = "http://localhost:8080/api"
        const respone = await fetch(url)
        const data = await respone.json()
        setText(JSON.stringify(data))
    }
    
    const [text, setText] = useState(() => "no resp yet...")

    return( 
    <div>
        <button onClick={fetchApi}>tap</button>
        <label id="response_label">{text}</label>
    </div>
    )
}

export default Test