import ShowRoomCode from "../pages/ShowRoomCode";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("ShowRoomCode", () => {
  it("Check render show room code", () => {
    render(<ShowRoomCode />);
    // check if all components are rendered
    expect(screen.getByTestId("buttonShowRoomCode")).toBeInTheDocument();
  });
});