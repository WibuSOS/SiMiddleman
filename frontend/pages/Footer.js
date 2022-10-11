const Footer = () => {
  return (
    <div>
      <footer>
        <div className="container">
          <div className="row d-flex justify-content-between">
            <div className="col-lg-4 col-sm-12 text-white" style={{margin: 40}}> 
              &#169; 2022, Made with Team 2 for a better web.
            </div>
            <div className="col-lg-4 col-sm-12 text-white" style={{margin: 40}}>
              <div className="row d-flex justify-content-between">
                <div className="col text-center">
                  <a href="/" className="link">Privacy Policy</a>
                </div>
                <div className="col text-center">
                  <a href="/" className="link">Terms &amp; Conditions</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default Footer