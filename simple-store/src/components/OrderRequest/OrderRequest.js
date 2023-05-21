import { useState } from "react";
import { Button, Form } from "react-bootstrap";

const CallRequestForm = () => {
  const [phone, setPhone] = useState("");
  const [information, setInformation] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSuccess, setIsSuccess] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsSubmitting(true);
    try {
      const response = await fetch("/api/request/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ phone, information }),
      });
      if (!response.ok) {
        throw new Error("Request failed");
      }
      setIsSuccess(true);
    } catch (error) {
      console.error(error);A
      setIsSubmitting(false);
    }
  };

  if (isSuccess) {
    return (
      <div className="text-center text-success mt-5">
        <p>Спасибо за обращение! Мы вам скоро перезвоним.</p>
      </div>
    );
  }

  return (
    <div className="d-flex justify-content-center align-items-center mt-5">
      <Form onSubmit={handleSubmit} className="w-25">
        <Form.Group controlId="phone">
          <Form.Label>Телефон</Form.Label>
          <Form.Control
            type="tel"
            placeholder="Введите телефон"
            value={phone}
            onChange={(event) => setPhone(event.target.value)}
            required
          />
        </Form.Group>
        <Form.Group controlId="information">
          <Form.Label>Описание</Form.Label>
          <Form.Control
            as="textarea"
            rows={3}
            placeholder="Введите описание"
            value={information}
            onChange={(event) => setInformation(event.target.value)}
          />
        </Form.Group>
        <Button className="mt-3" variant="success" type="submit" disabled={isSubmitting}>
          {isSubmitting ? "Отправка..." : "Заказать звонок"}
        </Button>
      </Form>
    </div>
  );
};

export default CallRequestForm;