import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { useGno } from "./api/use-gno";

function App() {
    const [resultText, setResultText] = useState(
        "Start GnoNative Kit on Wails to see result"
    );

    const gno = useGno();

    const onAction1 = () => {
        gno.listKeyInfo().then((res) => setResultText(res.toString()));
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div className="input-box"></div>
            <div id="result" className="result">
                {resultText}
            </div>
            <button className="btn" onClick={onAction1}>Action 1</button>
        </div>
    );
}

export default App;
