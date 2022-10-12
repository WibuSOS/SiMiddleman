import ModalShowRoomCode from "../pages/ModalShowRoomCode";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Modal Create Room", () => {
  it("Check ModalShowRoomCode.js", () => {
    render(<ModalShowRoomCode showRoomCodeModal={() => {}} />);
    // check if all components are rendered
    expect(screen.getByTestId("ModalHeader")).toBeInTheDocument();
    expect(screen.getByTestId("avatar")).toBeInTheDocument();
    expect(screen.getByTestId("logo")).toBeInTheDocument();
    expect(screen.getByTestId("title")).toBeInTheDocument();
    expect(screen.getByTestId("roomCode")).toBeInTheDocument();
    expect(screen.getByTestId("buttonSalin")).toBeInTheDocument();
  });
});