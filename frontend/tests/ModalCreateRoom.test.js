import ModalCreateRoom from "../pages/ModalCreateRoom";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Modal Create Room", () => {
  it("Check ModalCreateRoom.js", () => {
    render(<ModalCreateRoom closeCreateRoomModal={() => {}} createRoomModal={() => {}} />);
    // check if all components are rendered
    expect(screen.getByTestId("ModalCreateRoom")).toBeInTheDocument();
    expect(screen.getByTestId("ModalHeader")).toBeInTheDocument();
    expect(screen.getByTestId("avatar")).toBeInTheDocument();
    expect(screen.getByTestId("logo")).toBeInTheDocument();
    expect(screen.getByTestId("title")).toBeInTheDocument();
    expect(screen.getByTestId("namaProduk")).toBeInTheDocument();
    expect(screen.getByTestId("hargaProduk")).toBeInTheDocument();
    expect(screen.getByTestId("kuantitasProduk")).toBeInTheDocument();
    expect(screen.getByTestId("deskripsiProduk")).toBeInTheDocument();
    expect(screen.getByTestId("buttonCreateRoom")).toBeInTheDocument();
  });
});