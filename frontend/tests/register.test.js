import RegisterForm from "../pages/register";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Register", () => {
    it("Check register js", () => {
      render(<RegisterForm />);
      // check if all components are rendered
      expect(screen.getByText("Register")).toBeInTheDocument();
      expect(screen.getByTestId("buttonRegister")).toBeInTheDocument();
    });
});