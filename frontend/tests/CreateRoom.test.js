import CreateRoom from "../pages/CreateRoom";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Register", () => {
  it("Check register js", () => {
    render(<CreateRoom />);
    // check if all components are rendered
    expect(screen.getByTestId("createRoomButton")).toBeInTheDocument();
  });
});