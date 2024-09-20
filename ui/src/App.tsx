import { useState, useMemo } from "react";
import { FaClipboard } from "react-icons/fa";
import { IoSend, IoCheckmarkCircle } from "react-icons/io5";
import "./App.css";


type ColorMapping = {
  [key: number]: string;
};

function App() {

  const [duration, setDuration] = useState(-1);
  const [maxLength, setMaxLength] = useState(8);
  const [noise, setNoise] = useState(0);
  const [sentence, setSentence] = useState("");
  const [password, setPassword] = useState("");
  const [isVisible, setIsVisible] = useState(false);

  const getRandomColor = () => {
    // Generate a random hex color with maximum brightness
    const randomHexColor = () => {
      // Generate a random color component value from 0x00 to 0xFF
      const minValue = 100; // Change this value to adjust the lightness
      const r = Math.floor(Math.random() * (255 - minValue) + minValue);
      const g = Math.floor(Math.random() * (255 - minValue) + minValue);
      const b = Math.floor(Math.random() * (255 - minValue) + minValue);

      // Convert RGB to hex format
      return `#${r.toString(16).padStart(2, "0")}${g.toString(16).padStart(2, "0")}${b.toString(16).padStart(2, "0")}`;
    };

    return randomHexColor();
  };

  const capitalizeFirstWord = (sentence: string): string => {
    if (sentence.length === 0) return sentence; // Return early if the sentence is empty

    const [firstWord, ...restOfSentence] = sentence.split(/\s+/);
    const camelFirstWord = firstWord.charAt(0).toUpperCase() + firstWord.slice(1).toLowerCase();
    return [camelFirstWord, " ", ...restOfSentence.join(" ")].join("");
  };


  const createColorMapping = (sentence: string): ColorMapping => {
    const colorMapping: ColorMapping = {};
    const words = sentence.split(/\s+/); // Split by spaces to get words
    const letters = words.map(word => word.charAt(0)); // Extract the first letter of each word
    letters.forEach((_, index) => {
      if (!colorMapping[index]) {
        colorMapping[index] = getRandomColor();
      }
    });
    return colorMapping;
  };

  const createColoredSentence = (sentence: string) => {
    const words = sentence.split(/\s+/);
    return words.map((part, index) => {
      if (part.length > 0) {
        const firstLetter = part[0];
        const color = colorMapping[index] || "#f0f0f0"; // Default to black if no color found
        return (
          <span className="word" key={index}>
            <span style={{ color, fontWeight: "bold" }}>{firstLetter}</span>
            {part.slice(1)}
          </span>
        );
      }
      return part;
    });
  };


  const createPasswordDisplay = (password: string) => {

    let colorIndex = 0;

    return password.split("").map((letter, index) => {
      const isLetter = /[a-zA-Z]/.test(letter);
      const color = isLetter ? colorMapping[colorIndex++] || "#f0f0f0" : "#f0f0f0"; // Default to black if no color found
      return (
        <span key={index} style={{ color, fontWeight: isLetter ? "bold" : "normal" }}>
          {letter}
        </span >
      );
    });
  };

  const showNotification = () => {
    setIsVisible(true);
    setTimeout(() => {
      setIsVisible(false);
    }, 2000); // Notification will fade out after 2 seconds
  };

  const handleClipboard = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      showNotification();
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
  };

  const handleSubmit = async () => {

    try {
      const startTime = performance.now();
      const response = await fetch(`https://akro.fly.dev/generate?minLength=1&maxLength=${maxLength}&noiseLevel=${noise}`)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      const endTime = performance.now();
      setDuration(endTime - startTime);
      setSentence(capitalizeFirstWord(data.sentence));
      setPassword(data.password);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  const colorMapping = useMemo(() => createColorMapping(sentence), [sentence]);

  return (

    <div className="panel">
      <div className={`notification ${isVisible ? 'show' : 'hidden'}`}>
        <IoCheckmarkCircle /> <p>Copied to clipboard!</p>
      </div>
      <div className="title">
        <h1>akro</h1>
        <p>Create passwords with "slightly coherent" sentences. <br /> Built with <a href="https://simple.wikipedia.org/wiki/Markov_chain">Markov chains</a>, see the <a href="https://github.com/sweeneyngo/akro">details</a> + <a href="https://github.com/sweeneyngo/akro">code</a>.</p>
      </div>
      {duration >= 0 && <div className={`time ${duration >= 100 ? "time-red" : duration >= 48 && "time-yellow"}`}>
        <p>Returned {sentence.split(" ").filter(word => word.length > 0).length} word(s) in {duration.toFixed(2)}ms</p>
      </div>}
      <div className="sentence">
        <div className="output-container">
          <p className="output-text">{createColoredSentence(sentence)}</p>
          <div onClick={() => handleClipboard(sentence)} className="output-icon"><FaClipboard /></div>
          <p className="output-stats">{sentence.length} chars</p>
        </div>
      </div>
      <div className="password">
        <div className="output-container">
          <p className="output-text">{createPasswordDisplay(password)}</p>
          <div onClick={() => handleClipboard(password)} className="output-icon"><FaClipboard /></div>
          <p className="output-stats">{password.length} chars</p>
        </div>
      </div>
      <div className="flex-h">
        <div className="slider-container">
          <input
            type="range"
            aria-label="Range"
            min={1}
            max={20}
            step={1}
            value={maxLength}
            onChange={(e) => setMaxLength(parseInt(e.target.value))}
            className="slider"
          />
          <div className="slider-value"><p>Max length • {maxLength}</p></div>
        </div>
        <div className="slider-container">
          <input
            type="range"
            aria-label="Noise"
            min={0}
            max={20}
            step={1}
            value={noise}
            onChange={(e) => setNoise(parseInt(e.target.value))}
            className="slider"
          />
          <div className="slider-value"><p>Noise • {noise}</p></div>
        </div>
        <div className="button-container">
          <button id="submitButton" aria-label="Submit Button" onClick={() => handleSubmit()}>
            <div className="center">
              <IoSend />
            </div>
          </button>
        </div>

      </div>
      <div className="footer"><p>Made with love by <a href="https://www.ifuxyl.dev/">ifu</a></p></div>
    </div>

  )
}

export default App
