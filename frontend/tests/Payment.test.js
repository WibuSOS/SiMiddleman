import PaymentForm from "../pages/Payment";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

jest.mock('next/router', () => ({
  useRouter() {
    return ({
      route: '/',
      pathname: '',
      query: '',
      asPath: '',
      push: jest.fn(),
      events: {
        on: jest.fn(),
        off: jest.fn()
      },
      beforePopState: jest.fn(() => null),
      prefetch: jest.fn(() => null)
    });
  },
}));

describe("Payment", () => {
  it("Check Payment.js", () => {
    render(<PaymentForm />);
    // check if all components are rendered
    expect(screen.getByTestId("back_button")).toBeInTheDocument();
    expect(screen.getByTestId("sinarmas")).toBeInTheDocument();
    expect(screen.getByTestId("no_rek")).toBeInTheDocument();
    expect(screen.getByTestId("simiddleman")).toBeInTheDocument();
    expect(screen.getByTestId("harga")).toBeInTheDocument();
    expect(screen.getByTestId("instruction_no_rek")).toBeInTheDocument();
    expect(screen.getByTestId("question_payment")).toBeInTheDocument();
    expect(screen.getByTestId("instruction_receipt")).toBeInTheDocument();
    expect(screen.getByTestId("upload_receipt_button")).toBeInTheDocument();
  });
});
