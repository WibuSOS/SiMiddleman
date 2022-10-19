import { useState } from "react";
import NavItem from "./NavigationItem";
import Link from "next/link";
import Logo from './assets/logo2.png';
import { useRouter } from "next/router";
import { Dropdown } from "react-bootstrap";
import usFlag from './assets/usa.png';
import idFlag from "./assets/indonesia.png";

const MENU_LIST = [
  { text: "Home", href: "/" },
];

const Navbar = () => {
  const [navActive, setNavActive] = useState(null);
  const [activeIdx, setActiveIdx] = useState(-1);
  const router = useRouter();
  const pilihanBahasa = router.locales;

  const setFlagAndName = ( locale ) => {
    if (locale === "en") return (
      <a><img src={usFlag.src} className="logoBendera"/> English (USA)</a>
    )
    else return (
      <a><img src={idFlag.src} className="logoBendera"/> Indonesia</a>
    )
  }

  return (
    <header>
      <nav className='container nav'>
        <Link href={"/"}>
          <a className="logo-brand">
            <img src={Logo.src} className='logo-navbar me-3'/>
            SiMiddleman+
          </a>
        </Link>
        <div onClick={() => setNavActive(!navActive)} className='nav__menu-bar'>
          <div></div>
          <div></div>
          <div></div>
        </div>
        <div className={`${navActive ? "active" : ""} nav__menu-list`}>
          <Dropdown>
            <Dropdown.Toggle variant="simiddleman" id="dropdown-basic">
            {setFlagAndName(router.locale)}
            </Dropdown.Toggle>
            <Dropdown.Menu>
              {pilihanBahasa.map((locale, key) => (
                <Dropdown.Item key={key}>
                  <Link href={router.asPath} locale={locale}>
                  {setFlagAndName(locale)}
                  </Link>
                </Dropdown.Item>
              ))}
            </Dropdown.Menu>
          </Dropdown>
          {/* {MENU_LIST.map((menu, idx) => {
            return (
              <div className="link-div" key={menu.text} onClick={() => {
                setActiveIdx(idx);
                setNavActive(false);
              }}>
                <NavItem active={activeIdx === idx} {...menu} />
              </div>
            )
          })} */}
        </div>
      </nav>
    </header>
  );
};

export default Navbar;