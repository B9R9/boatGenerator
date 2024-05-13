const DisplayBoatsDetails = ({ boat, handleClick, handleHistory }) => {

	return (
		<>
		<div style={{
              position: 'absolute',
			  top: '90%',
			  left: '50%',
			  transform: 'translate(-50%, -50%)',
			  zIndex: '9999',
			  backgroundColor: 'rgba(173, 216, 230, 0.7)',
			  borderRadius: '100px',
			  padding: '20px',
			  width: '50vw',
			  height: '5vh',
			  display: 'flex',
			  flexDirection: 'row', // Aligner les éléments horizontalement
			  justifyContent: 'space-between',
			  alignItems: 'center', // Aligner les éléments verticalement au centre

			}}>
		<button>{"<"}</button>
		<div style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-between', gap: '20px', fontFamily: 'Arial'}}>
        <p style={{ paddingLeft: '5px' }}>ID: {boat.id}</p>
        <p>Speed: {boat.Speed}</p>
        <p>Latitude: {boat.Latitude}</p>
        <p>Longitude: {boat.Longitude}</p>
    </div>
		<button>{">"}</button>
		<button onClick={handleClick}>Close</button>
		{/* <button onClick={handleHistory}>Show trajectoire</button> */}
		</div>
		</>
	)
}

export default DisplayBoatsDetails