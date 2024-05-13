import { useState } from 'react'

const DisplaySettings = ({ onChange }) => {
	const [selectedOption, setSelectedOption] = useState(1);
	const [customValue, setCustomValue] = useState('');
  
	const handleSelectChange = (e) => {
	  const value = parseInt(e.target.value);
	  setSelectedOption(value);
	  setCustomValue('');
	  onChange(value * 1000);
	};
  
	const handleCustomInputChange = (e) => {
	  const value = e.target.value;
	  setSelectedOption(-1); // -1 pour indiquer une valeur personnalisée
	  setCustomValue(value);
	  onChange(parseInt(value) * 1000);
	};
  
	return (
	  <div>
		<select value={selectedOption} onChange={handleSelectChange}>
		  {[...Array(10).keys()].map((index) => (
			<option key={index + 1} value={index + 1}>{index + 1}</option>
		  ))}
		  <option value="-1">Autre...</option>
		</select>
		{selectedOption === -1 && (
		  <input type="number" value={customValue} onChange={handleCustomInputChange} />
		)}
	  </div>
	)
  }
export default DisplaySettings