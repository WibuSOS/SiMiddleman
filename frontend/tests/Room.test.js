import RoomChat from "../pages/Room";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

describe("Room Chat", () => {
  it("Check render room", () => {
    render(<RoomChat />);
    // check if all components are rendered
    expect(screen.getByTestId("container")).toBeInTheDocument();
    expect(screen.getByTestId("Close")).toBeInTheDocument();
    expect(screen.getByTestId("container2")).toBeInTheDocument();
    expect(screen.getByTestId("title")).toBeInTheDocument();
    expect(screen.getByTestId("subTitle")).toBeInTheDocument();
    expect(screen.getByTestId("buttonEdit")).toBeInTheDocument();
    expect(screen.getByTestId("titleNamaProduk")).toBeInTheDocument();
    expect(screen.getByTestId("namaProduk")).toBeInTheDocument();
    expect(screen.getByTestId("titleKuantitas")).toBeInTheDocument();
    expect(screen.getByTestId("kuantitasProduk")).toBeInTheDocument();
    expect(screen.getByTestId("titleHarga")).toBeInTheDocument();
    expect(screen.getByTestId("hargaProduk")).toBeInTheDocument();
    expect(screen.getByTestId("titleDeskripsi")).toBeInTheDocument();
    expect(screen.getByTestId("deskripsiProduk")).toBeInTheDocument();
    expect(screen.getByTestId("buttonCheckout")).toBeInTheDocument();
  });
});