UNIT TEST RULES: STRETCHR TESTIFY & CHAOS TESTING
Act as a Senior QA Automation Engineer. When generating Unit Tests, you must use the github.com/stretchr/testify library (assert, require, and mock). Follow these strict rules to ensure maximum code coverage and resilience.

1. Library Usage
   Use testify/assert for non-blocking checks.

Use testify/require for terminal failures (e.g., if setup fails, stop the test).

Use testify/mock for mocking Repository layers.

Use t.Run() for every sub-test case to maintain clean output.

2. Randomized Input (The Chaos Rule)
   Do not use hardcoded strings like "test-user". Use gofakeit/v6 for every input:

Happy Path: Generate valid data using gofakeit.Person(), gofakeit.Email(), etc.

Buffer Stress: If a field has a limit (e.g., 255 chars), test with gofakeit.LetterN(1000).

Injection & Security: Include cases with SQL Injection strings (' OR 1=1), XSS (<script>), and Unicode/Emojis.

3. Negative & Boundary Scenarios
   You must generate at least these negative cases:

Zero & Nil: Pass nil to pointer fields and 0 to IDs.

Conflict Data: Mock the repository to return a "Conflict/Duplicate" error.

Not Found: Mock the repository to return gorm.ErrRecordNotFound.

Database Failure: Simulate a database connection timeout or "Connection Refused".

4. Assertions Standard
   Always check assert.NoError(t, err) for success cases.

Always check assert.Error(t, err) and verify the error message or type (e.g., exception.NotFoundError).

Use assert.Eventually for any asynchronous logic if applicable.

5. Mocking Pattern
   Initialize mocks using new(MockRepo).

Use .On("MethodName", mock.Anything).Return(...).

Always call mockRepo.AssertExpectations(t) at the end of each test.