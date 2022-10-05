import RegisterModal from "../pages/registerModal";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Register Form", () => {
  it("Renders a register form", () => {
    render(<RegisterModal handleSubmitRegister={() => {}}
      closeRegisterModal={() => {}}
      registerModal={() => {}}/>);
    // check if all components are rendered
    expect(screen.getByTestId("nama")).toBeInTheDocument();
    expect(screen.getByTestId("noHp")).toBeInTheDocument();
    expect(screen.getByTestId("noRek")).toBeInTheDocument();
    expect(screen.getByTestId("email")).toBeInTheDocument();
    expect(screen.getByTestId("password")).toBeInTheDocument();
    expect(screen.getByTestId("confirm")).toBeInTheDocument();
    expect(screen.getByTestId("modalRegister")).toBeInTheDocument();
    expect(screen.getByTestId("modalHeader")).toBeInTheDocument();
    expect(screen.getByTestId("logo")).toBeInTheDocument();
    expect(screen.getByTestId("title")).toBeInTheDocument();
    expect(screen.getByTestId("modalBody")).toBeInTheDocument();
    expect(screen.getByTestId("form")).toBeInTheDocument();
    expect(screen.getByTestId("submitButton")).toBeInTheDocument();
  });
});