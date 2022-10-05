import LoginForm from "../pages/Login";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Login", () => {
    it("Renders a login form", () => {
      render(<LoginForm />);
      // check if all components are rendered
      expect(screen.getByTestId("main")).toBeInTheDocument();
      expect(screen.getByTestId("centeredForm")).toBeInTheDocument();
      expect(screen.getByTestId("avatar")).toBeInTheDocument();
      expect(screen.getByTestId("logo")).toBeInTheDocument();
      expect(screen.getByTestId("Title")).toBeInTheDocument();
      expect(screen.getByTestId("loginForm")).toBeInTheDocument();
      expect(screen.getByTestId("email")).toBeInTheDocument();
      expect(screen.getByTestId("password")).toBeInTheDocument();
      expect(screen.getByTestId("LupaPassword")).toBeInTheDocument();
      expect(screen.getByTestId("masukButton")).toBeInTheDocument();
      expect(screen.getByTestId("OR")).toBeInTheDocument();
      expect(screen.getByTestId("buttonDaftar")).toBeInTheDocument();
    });
});