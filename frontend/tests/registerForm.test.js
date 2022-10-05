import RegForm from "../pages/registerForm";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Register Form", () => {
    it("renders a register form", () => {
      render(<RegForm handleSubmitRegister={() => {}} />);
      // check if all components are rendered
      expect(screen.getByTestId("nama")).toBeInTheDocument();
      expect(screen.getByTestId("noHp")).toBeInTheDocument();
      expect(screen.getByTestId("noRek")).toBeInTheDocument();
      expect(screen.getByTestId("email")).toBeInTheDocument();
      expect(screen.getByTestId("password")).toBeInTheDocument();
      expect(screen.getByTestId("confirmPassword")).toBeInTheDocument();
    });
  });