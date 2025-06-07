package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type IRVisitor struct {
	*parser.BaseKumirParserVisitor
}

func NewIRVisitor() *IRVisitor {
	return &IRVisitor{
		BaseKumirParserVisitor: &parser.BaseKumirParserVisitor{},
	}
}

func (v *IRVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *IRVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("visit Program")
	for _, item := range ctx.AllProgramItem() {
		v.Visit(item)
	}

	for _, item := range ctx.AllAlgorithmDefinition() {
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitProgramItem(ctx *parser.ProgramItemContext) interface{} {
	fmt.Println("visit ProgramItem")

	switch {
	case ctx.GlobalDeclaration() != nil:
		v.Visit(ctx.GlobalDeclaration())
	case ctx.GlobalAssignment() != nil:
		v.Visit(ctx.GlobalAssignment())
	default:
		return nil
	}

	return nil
}

func (v *IRVisitor) VisitGlobalDeclaration(ctx *parser.GlobalDeclarationContext) interface{} {
	fmt.Println("visit GlobalDeclaration")
	return nil
}

func (v *IRVisitor) VisitGlobalAssignment(ctx *parser.GlobalAssignmentContext) interface{} {
	fmt.Println("visit GlobalAssignment")
	return nil
}

func (v *IRVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) interface{} {
	fmt.Println("visit AlgorithmDefinition")
	v.Visit(ctx.AlgorithmHeader())
	for _, d := range ctx.AllVariableDeclaration() {
		v.Visit(d)
	}
	for _, p := range ctx.AllPreCondition() {
		v.Visit(p)
	}
	for _, p := range ctx.AllPostCondition() {
		v.Visit(p)
	}
	v.Visit(ctx.AlgorithmBody())

	return nil
}

func (v *IRVisitor) VisitAlgorithmHeader(ctx *parser.AlgorithmHeaderContext) interface{} {
	fmt.Println("visit AlgorithmHeader")
	return nil
}

func (v *IRVisitor) VisitAlgorithmBody(ctx *parser.AlgorithmBodyContext) interface{} {
	fmt.Println("visit AlgorithmBody")
	v.Visit(ctx.StatementSequence())
	return nil
}

func (v *IRVisitor) VisitPreCondition(ctx *parser.PreConditionContext) interface{} {
	fmt.Println("visit PreCondition")
	return nil
}

func (v *IRVisitor) VisitPostCondition(ctx *parser.PostConditionContext) interface{} {
	fmt.Println("visit PostCondition")
	return nil
}

func (v *IRVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	fmt.Println("visit VariableDeclaration")
	return nil
}

func (v *IRVisitor) VisitStatementSequence(ctx *parser.StatementSequenceContext) interface{} {
	fmt.Println("visit StatementSequence")
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	return nil
}

func (v *IRVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("visit Statement")

	switch {
	case ctx.VariableDeclaration() != nil:
		v.Visit(ctx.VariableDeclaration())
	case ctx.AssignmentStatement() != nil:
		v.Visit(ctx.AssignmentStatement())
	case ctx.IoStatement() != nil:
		v.Visit(ctx.IoStatement())
	case ctx.IfStatement() != nil:
		v.Visit(ctx.IfStatement())
	case ctx.SwitchStatement() != nil:
		v.Visit(ctx.SwitchStatement())
	case ctx.LoopStatement() != nil:
		v.Visit(ctx.LoopStatement())
	case ctx.ExitStatement() != nil:
		v.Visit(ctx.ExitStatement())
	case ctx.PauseStatement() != nil:
		v.Visit(ctx.PauseStatement())
	case ctx.StopStatement() != nil:
		v.Visit(ctx.StopStatement())
	//case ctx.AssertionStatement() != nil:
	//	v.Visit(ctx.AssertionStatement())
	default:
		return nil
	}

	return nil
}

func (v *IRVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	fmt.Println("visit AssignmentStatement")

	if ctx.Lvalue() != nil {
		v.Visit(ctx.Lvalue())
	}

	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
	}

	return nil
}

func (v *IRVisitor) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	fmt.Println("visit Lvalue")
	if ctx.QualifiedIdentifier() != nil {
		v.Visit(ctx.QualifiedIdentifier())
	}

	return nil
}

func (v *IRVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	fmt.Println("visit Expression")

	if ctx.LogicalOrExpression() != nil {
		v.Visit(ctx.LogicalOrExpression())
	}

	return nil
}

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	fmt.Println("visit Literal")
	if ctx.INTEGER() != nil {
		fmt.Println("visit INTEGER", ctx.INTEGER().GetText())
		return nil
	}

	if ctx.REAL() != nil {
		fmt.Println("visit REAL", ctx.REAL().GetText())
		return nil
	}

	if ctx.STRING() != nil {
		fmt.Println("visit STRING", ctx.STRING().GetText())
		return nil
	}

	if ctx.CHAR_LITERAL() != nil {
		fmt.Println("visit CHAR_LITERAL", ctx.CHAR_LITERAL().GetText())
		return nil
	}

	if ctx.TRUE() != nil {
		fmt.Println("visit TRUE", ctx.TRUE().GetText())
		return nil
	}

	if ctx.FALSE() != nil {
		fmt.Println("visit FALSE", ctx.FALSE().GetText())
		return nil
	}

	if ctx.NEWLINE_CONST() != nil {
		fmt.Println("visit NEWLINE_CONST", ctx.NEWLINE_CONST().GetText())
		return nil
	}

	return nil
}

func (v *IRVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	fmt.Println("visit PrimaryExpression")
	if ctx.Literal() != nil {
		v.Visit(ctx.Literal())
	}
	return nil
}

func (v *IRVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	fmt.Println("visit PostfixExpression")

	if ctx.PrimaryExpression() != nil {
		v.Visit(ctx.PrimaryExpression())
	}

	return nil
}

func (v *IRVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	fmt.Println("visit UnaryExpression")

	operator := ctx.GetOp()
	if operator == nil {
		if ctx.PostfixExpression() != nil {
			v.Visit(ctx.PostfixExpression())
		}
		return nil
	}
	operatorTokenType := operator.GetTokenType()
	switch operatorTokenType {
	case parser.KumirParserPLUS:
		fmt.Println("visit KumirParserPLUS")
	case parser.KumirParserMINUS:
		fmt.Println("visit KumirParserMINUS")
	case parser.KumirParserNOT:
		fmt.Println("visit KumirParserNOT")
	default:
		return nil
	}

	if ctx.UnaryExpression() != nil {
		v.Visit(ctx.UnaryExpression())
	}

	return nil
}

func (v *IRVisitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	fmt.Println("visit PowerExpression")

	if ctx.UnaryExpression() != nil {
		v.Visit(ctx.UnaryExpression())
	}

	return nil
}

func (v *IRVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	fmt.Println("visit MultiplicativeExpression")
	for _, item := range ctx.AllPowerExpression() {
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	fmt.Println("visit AdditiveExpression")

	for _, item := range ctx.AllMultiplicativeExpression() {
		//spew.Dump(item)
		v.Visit(item)
	}

	if len(ctx.AllPLUS()) > 0 {
		fmt.Println("slozit'", len(ctx.AllPLUS()))
		return nil
	}

	//switch ctx.GetOperation().GetTokenType() {
	//case antlr.:
	//
	//}

	return nil
}

func (v *IRVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	fmt.Println("visit RelationalExpression")
	for _, item := range ctx.AllAdditiveExpression() {
		v.Visit(item)
	}
	return nil
}

func (v *IRVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	fmt.Println("visit EqualityExpression")
	for _, item := range ctx.AllRelationalExpression() {
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	fmt.Println("visit LogicalAndExpression")
	for _, item := range ctx.AllEqualityExpression() {
		v.Visit(item)
	}
	return nil
}

func (v *IRVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	fmt.Println("visit LogicalOrExpression")

	for _, item := range ctx.AllLogicalAndExpression() {
		//spew.Dump(item)
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitIoStatement(ctx *parser.IoStatementContext) interface{} {
	fmt.Println("visit IoStatement")
	return nil
}

func (v *IRVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	fmt.Println("visit IfStatement")
	return nil
}

func (v *IRVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	fmt.Println("visit SwitchStatement")
	return nil
}

func (v *IRVisitor) VisitLoopStatement(ctx *parser.LoopStatementContext) interface{} {
	fmt.Println("visit LoopStatement")
	v.Visit(ctx.StatementSequence())
	return nil
}

func (v *IRVisitor) VisitExitStatement(ctx *parser.ExitStatementContext) interface{} {
	fmt.Println("visit ExitStatement")
	return nil
}

func (v *IRVisitor) VisitPauseStatement(ctx *parser.PauseStatementContext) interface{} {
	fmt.Println("visit PauseStatement")
	return nil
}

func (v *IRVisitor) VisitStopStatement(ctx *parser.StopStatementContext) interface{} {
	fmt.Println("visit StopStatement")
	return nil
}

func (v *IRVisitor) VisitQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) interface{} {
	fmt.Println("visit QualifiedIdentifier")
	return nil
}
