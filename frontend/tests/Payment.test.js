import PaymentForm from "../pages/Payment";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

// jest.mock('next/router', () => ({
//   useRouter() {
//     return ({
//       pathname: '',
//       query: {},
//       asPath: '',
//       push: jest.fn(),
//       events: {
//         on: jest.fn(),
//         off: jest.fn()
//       },
//       beforePopState: jest.fn(() => null),
//       prefetch: jest.fn(() => null)
//     });
//   },
// }));

describe("when clicking the beli button", () => {
  // let component: ReactWrapper;

  // beforeEach(async () => {
  //   const useRouter = jest.spyOn(require("next/router"), "useRouter");

  //   useRouter.mockImplementation(() => ({
  //     pathname: '/Payment',
  //     query: {
  //       idRoom: `${1}`,
  //     },
  //     asPath: '/Payment',
  //     push: jest.fn(),
  //     events: {
  //       on: jest.fn(),
  //       off: jest.fn()
  //     },
  //     beforePopState: jest.fn(() => null),
  //     prefetch: jest.fn(() => null)
  //   }));

  //   await waitFor(() => {
  //     component = mount(
  //       <Provider store={store}>
  //         <FavoritesPage />
  //       </Provider>
  //     );
  //   });
  //   component.find('.price-drop-nav-item').simulate('click');
  // });

  it("check Payment.js", () => {
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
