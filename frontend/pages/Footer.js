import Link from "next/link"

const Footer = () => {
  return (
    <div>
      <footer>
        <div className="container">
          <div className="footer-content row d-flex justify-content-between">
            <div className="col-lg-4 col-md-6 col-sm-6 text-white text-center"> 
              &#169; 2022, Made with Team 2 for a better web.
            </div>
            <div className="col-lg-4 col-md-6 col-sm-6 text-white">
              <div className="row d-flex justify-content-between link">
                <div className="col-sm-6 text-center">
                  <Link href="/">Privacy Policy</Link>
                </div>
                <div className="col-sm-6 text-center">
                  <Link href="/">Terms &amp; Conditions</Link>
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